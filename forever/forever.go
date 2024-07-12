package forever

import (
	"forever/task"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

type Forever struct {
	config *task.Config
	tasks  []*task.Task
}

func (forever *Forever) Start() {
	for _, t := range forever.tasks {
		err := t.Start(task.ForeverLifeCycleStart)
		if err != nil {
			logrus.Warnf("execute task err:%v", err)
		}
	}

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case aSignal := <-signals:
		if aSignal == syscall.SIGINT || aSignal == syscall.SIGTERM {
			logrus.Infof("Received signal: %v", aSignal)
			for _, t := range forever.tasks {
				t.Stop()
			}
		}
	}
}

func (forever *Forever) Close() {

}

func NewRunner(config *task.Config) (*Forever, error) {
	tasks := make([]*task.Task, len(config.Tasks))
	for i, t := range config.Tasks {
		t, err := task.NewTask(t)
		if err != nil {
			logrus.Warnf("create task err:%v", err)
		}
		tasks[i] = t
	}
	return &Forever{config, tasks}, nil
}
