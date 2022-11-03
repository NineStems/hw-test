package config

import (
	"os"
	"time"

	"gopkg.in/yaml.v2" //nolint

	"github.com/hw-test/hw12_13_14_15_calendar/pkg/errors" //nolint
)

// Logger конфигурация для логгера.
type Logger struct {
	Path  string `yaml:"path"`
	Level string `yaml:"level"`
}

type Grpc struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Rest struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Server конфигурация для HTTP сервера.
type Server struct {
	Grpc Grpc `yaml:"grpc"`
	Http Rest `yaml:"rest"`
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

type Scheduler struct {
	Pause time.Duration `yaml:"pause"`
}

// Rabbit конфигурация очереди.
type Rabbit struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Exchange   string `yaml:"exchange"`
	Queue      string `yaml:"queue"`
	Key        string `yaml:"key"`
	Credential struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"credential"`
}

// Config конфигурация сервиса.
type Config struct {
	Logger    Logger    `yaml:"logger"`
	Scheduler Scheduler `yaml:"scheduler"`
	Rabbit    Rabbit    `yaml:"rabbit"`
	Server    Server    `yaml:"server"`
	Database  Database  `yaml:"database"`
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
