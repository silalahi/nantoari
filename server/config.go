package server

import "fmt"

// Config is a struct to configure server
type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// Addr return HTTP address in string format
func (c Config) Addr() string {
	return fmt.Sprint(c.Host, ":", c.Port)
}
