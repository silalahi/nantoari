package store

import (
	"github.com/go-redis/redis"
)

// RedisStore is the Redis implementation of store.Interface.
type RedisStore struct {
	client     *redis.Client
	expiration int    // currently no use
	prefix     string // currently no use
}

// Get returns value for the given key.
func (rs *RedisStore) Get(key string) (interface{}, error) {
	return rs.client.Get(key).Result()
}

// Set returns value for the given key.
func (rs *RedisStore) Set(key string, value interface{}) error {
	return rs.client.Set(key, value, 0).Err()
}

// NewRedisStore returns Redis client instance of store.Interface.
func NewRedisStore(config *RedisConfig) (Interface, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     config.Addr(),
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := r.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisStore{
		client: r,
	}, nil
}
