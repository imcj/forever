package task

import (
	"bufio"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

type Task struct {
	Cmd        *exec.Cmd
	ProcessId  int
	Output     io.ReadCloser
	Error      io.ReadCloser
	outputFile *os.File
	errorFile  *os.File
}

func (task *Task) Execute() error {
	err := task.Cmd.Start()
	if err != nil {
		return err
	}
	go task.ReadOutput()
	go task.ReadError()
	task.ProcessId = task.Cmd.Process.Pid

	return nil
}

func (task *Task) ReadOutput() {
	scanner := bufio.NewScanner(task.Output)
	task.PipeToFile(scanner, task.outputFile)
}

func (task *Task) ReadError() {
	scanner := bufio.NewScanner(task.Error)
	task.PipeToFile(scanner, task.errorFile)
}

func (task *Task) PipeToFile(scanner *bufio.Scanner, file *os.File) {
	if file == nil {
		return
	}
	for scanner.Scan() {
		line := scanner.Text()
		logrus.Info(line)
		_, err := file.Write([]byte(line + "\n"))
		if err != nil {
			logrus.Warnf("Error writing to output file: %v", err)
		}
		err = file.Sync()
		if err != nil {
			logrus.Warnf("Error syncing to output file: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		logrus.Warnf("Error reading output: %v", err)
	}
}

func (task *Task) Kill() error {
	return task.Cmd.Process.Kill()
}

func OpenLoggingFile(path string) (*os.File, error) {
	if path == "" {
		return nil, errors.New("path is nil")
	}
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}

	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func NewTask(cmd *TaskConfig) (*Task, error) {
	c := cmd.Command.String()
	command := exec.Command(c, cmd.Arguments...)

	if cmd.Directory != "" {
		command.Dir = cmd.Directory.String()
	}

	output, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}

	standardError, err := command.StderrPipe()
	if err != nil {
		return nil, err
	}

	outputFile, err := OpenLoggingFile(cmd.OutputPath)
	errorFile, err := OpenLoggingFile(cmd.ErrorPath)

	return &Task{
		Cmd:        command,
		Output:     output,
		Error:      standardError,
		ProcessId:  0,
		outputFile: outputFile,
		errorFile:  errorFile,
	}, nil
}
