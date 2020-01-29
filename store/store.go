package store

import (
	"fmt"
)

const (
	// RedisStoreDriver is store driver using Redis
	RedisStoreDriver = "redis"
)

// Interface is the store interface.
type Interface interface {
	// Get returns value for the given key.
	Get(key string) (interface{}, error)

	// Set sets value for the given key.
	Set(key string, value interface{}) error

	// Exists checks if the given key exists.
	// Exists(key string) (bool, error)

	// Delete deletes the given key.
	// Delete(key string) error

	// Flush flushes the store.
	// Flush() error

	// Return all keys matching pattern
	// Keys(pattern string) ([]interface{}, error)

	// Close closes the connection to the store.
	// Close() error
}

// NewStore returns a store instance implementation from config
func NewStore(cfg *Config) (Interface, error) {
	switch cfg.Driver {
	case RedisStoreDriver:
		return NewRedisStore(&cfg.Redis)
	}

	return nil, fmt.Errorf("store driver %s does not exist", cfg.Driver)
}
