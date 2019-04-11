package config

// Config that the server needs
type Config struct {
	port string
}

// Loader loads the config
type Loader interface {
	Load() (*Config, error)
}
