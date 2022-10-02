package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2" // nolint:gci,typecheck

	"github.com/calendar/hw12_13_14_15_calendar/pkg/errors"
)

// Logger конфигурация для логгера.
type Logger struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

// Server конфигурация для HTTP сервера.
type Server struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

// Database конфигурация для базы данных.
type Database struct {
	Source   string        `yaml:"source"`
	Username string        `yaml:"user"`
	Password string        `yaml:"pass"`
	Host     string        `yaml:"host"`
	Port     int           `yaml:"port"`
	Database string        `yaml:"database"`
	Timeout  time.Duration `yaml:"timeout"`
}

// Config конфигурация сервиса.
type Config struct {
	Logger   Logger   `yaml:"logger"`
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

// Apply применяет значение из конфигурационного файла.
func (c *Config) Apply(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "os.Open")
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	if err = decoder.Decode(c); err != nil {
		return errors.Wrap(err, "decoder.Decode")
	}

	return nil
}

func New() *Config {
	return &Config{}
}
