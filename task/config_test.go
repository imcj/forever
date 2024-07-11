package task_test

import (
	"forever/task"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestConfig_LoadConfig(t *testing.T) {
	cwd, err := os.Getwd()
	assert.Nil(t, err)

	configFile := filepath.Join(cwd, "../", "forever.yml")
	config, err := task.LoadConfig(configFile)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(config.Tasks))
}
