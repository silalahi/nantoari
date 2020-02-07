package server

const (
	// DefaultServerPort is the default port of the application server
	DefaultServerPort = 8080
	// DefaultServerHost is the default host of the application server
	DefaultServerHost = "localhost"
)

// Config is a struct to configure server
type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

// DefaultConfig returns a default config instance
func DefaultConfig() *Config {
	return &Config{
		Host: DefaultServerHost,
		Port: DefaultServerPort,
	}
}
