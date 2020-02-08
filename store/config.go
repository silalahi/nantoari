package store

import (
	"fmt"
)

// Config is a struct to represent a key/value store
type Config struct {
	Prefix string      `yaml:"prefix"`
	Driver string      `yaml:"driver"`
	Redis  RedisConfig `yaml:"redis"`
}

// RedisConfig is a struct contain store configuration for Redis implementation
type RedisConfig struct {
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Password   string `yaml:"password"`
	DB         int    `yaml:"db"`
	Expiration int    `yaml:"expiration"`
}

// Addr returns Redis address in string format
func (r RedisConfig) Addr() string {
	return fmt.Sprint(r.Host, ":", r.Port)
}
