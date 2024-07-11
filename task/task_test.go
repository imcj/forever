package task

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestTask_Execute(t *testing.T) {
	workingDir, err := os.Getwd()
	workingDir = filepath.Join(workingDir, "../")

	t.Logf("working dir: %s", workingDir)
	assert.Nil(t, err)
	task, err := NewTask(&TaskConfig{
		Command:    Path(filepath.Join("./sleep/sleep")),
		Arguments:  []string{},
		Directory:  Path(filepath.Join(workingDir, "./")),
		OutputPath: filepath.Join(workingDir, "./logs/output.log"),
		ErrorPath:  filepath.Join(workingDir, "./logs/error.log"),
	})

	assert.Nil(t, err)

	go func() {
		err := task.Execute()
		if err != nil {
			panic(err)
		}
		t.Logf("Process id %d", task.ProcessId)
	}()

	time.Sleep(2 * time.Second)

	err = task.Kill()
	assert.Nil(t, err)
}
