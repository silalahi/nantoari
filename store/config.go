package store

import (
	"fmt"
)

// Config is a struct to represent a key/value store (Redis)
type Config struct {
	Prefix string      `yaml:"prefix"`
	Driver string      `yaml:"driver"`
	Redis  RedisConfig `yaml:"redis"`
}

// RedisConfig is a struct contain Redis configuration
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Addr returns Redis address in string format
func (rc RedisConfig) Addr() string {
	return fmt.Sprint(rc.Host, ":", rc.Port)
}
