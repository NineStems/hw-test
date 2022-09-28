package config

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/fixme_my_friend/hw12_13_14_15_calendar/pkg/errors"
)

type Config struct {
	Logger struct {
		Path  string `yaml:"path"`
		Level string `yaml:"level"`
	} `yaml:"logger"`
	Server struct {
		Port string `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

// Apply применяет значение из конфигурационного файла.
func (c *Config) Apply(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "os.Open")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(c)
	if err != nil {
		return errors.Wrap(err, "decoder.Decode")
	}

	return nil
}

func New() *Config {
	return &Config{}
}
