package config

// Config that the server needs
type Config struct {
	Port string
}

// Loader loads the config
type Loader interface {
	Load() (*Config, error)
}
