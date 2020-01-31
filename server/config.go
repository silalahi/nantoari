package server

const (
	// DefaultServerPort is the default port of the application server
	DefaultServerPort = 8080
)

// Config is a struct to configure server
type Config struct {
	Port int `yaml:"port"`
}
