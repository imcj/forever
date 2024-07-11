package task

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Config struct {
	Tasks []*TaskConfig `yaml:"tasks"`
}

type TaskConfig struct {
	Command    Path     `yaml:"command"`
	Arguments  []string `yaml:"arguments"`
	OutputPath string   `yaml:"output"`
	ErrorPath  string   `yaml:"error"`
	Directory  Path     `yaml:"directory"`
}

func LoadConfig(path string) (*Config, error) {
	config := &Config{}
	working, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	if !filepath.IsAbs(path) {
		path = filepath.Join(working, path)
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			logrus.Warnf("Error closing file: %v", err)
		}
	}(file)
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
