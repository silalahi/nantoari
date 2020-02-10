package store

import (
	"github.com/go-redis/redis"
)

// RedisStore is the Redis implementation of KV.
type RedisStore struct {
	client     *redis.Client
	expiration int
	prefix     string
}

// Get returns value for the given key.
func (r *RedisStore) Get(key string) (interface{}, error) {
	return r.client.Get(key).Result()
}

// Set returns value for the given key.
func (r *RedisStore) Set(key string, value interface{}) error {
	return r.client.Set(key, value, 0).Err()
}

// SetPrefix sets KV prefix
func (r *RedisStore) SetPrefix(prefix string) {
	r.prefix = prefix
}

// NewRedisStore returns Redis client instance of store.Interface.
func NewRedisStore(cfg *RedisConfig) (KV, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr(),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := r.Ping().Result()
	if err != nil {
		return nil, err
	}

	return &RedisStore{
		client:     r,
		expiration: cfg.Expiration,
		prefix:     "",
	}, nil
}
