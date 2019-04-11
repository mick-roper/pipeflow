package config

import "os"

// EnvLoader loads config from the environment
type EnvLoader struct {
}

// Load the config
func (l *EnvLoader) Load() (*Config, error) {
	port := os.Getenv("PORT")

	cfg := &Config{Port: port}

	return cfg, nil
}

// NewEnvLoader creates a new EnvLoader
func NewEnvLoader() *EnvLoader {
	return &EnvLoader{}
}
