package task

import "github.com/sirupsen/logrus"

type Runner struct {
	config *Config
	tasks  []*Task
}

func (runner *Runner) Run() {
	for _, task := range runner.tasks {
		err := task.Execute()
		if err != nil {
			logrus.Warnf("execute task err:%v", err)
		}
	}
	select {}
}

func (runner *Runner) Close() {

}

func NewRunner(config *Config) (*Runner, error) {
	tasks := make([]*Task, len(config.Tasks))
	for i, task := range config.Tasks {
		t, err := NewTask(task)
		if err != nil {
			logrus.Warnf("create task err:%v", err)
		}
		tasks[i] = t
	}
	return &Runner{config, tasks}, nil
}
