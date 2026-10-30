package logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type LoggerConfig struct {
	Level string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewLoggerConfig() (LoggerConfig, error) {
	var config LoggerConfig

	//"LOGGER" is a prefix to short naming of envconfig in struct
	if err := envconfig.Process("LOGGER", &config); err != nil {
		return LoggerConfig{}, fmt.Errorf("Process envconfig: %w", err)
	}

	return config, nil
}

func NewLoggerConfigMust() LoggerConfig {
	config, err := NewLoggerConfig()

	if err != nil {
		fmt.Errorf("Get Logger config: %w", err)
		panic(err)
	}

	return config
}