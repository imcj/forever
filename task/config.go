package task

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"io"
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
	Autostart  *bool    `yaml:"autostart"`
}

func LoadConfig(path string) (*Config, error) {

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

	return LoadConfigText(file)
}

func LoadConfigText(reader io.Reader) (*Config, error) {
	config := &Config{}
	decoder := yaml.NewDecoder(reader)
	err := decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
