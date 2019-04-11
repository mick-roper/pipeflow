package config

import (
	"errors"
	"os"
)

// EnvLoader loads config from the environment
type EnvLoader struct {
}

// Load the config
func (l *EnvLoader) Load() (*Config, error) {
	port := os.Getenv("PORT")

	if port == "" {
		return nil, errors.New("PORT has not been defined")
	}

	cfg := &Config{Port: port}

	return cfg, nil
}

// NewEnvLoader creates a new EnvLoader
func NewEnvLoader() *EnvLoader {
	return &EnvLoader{}
}
