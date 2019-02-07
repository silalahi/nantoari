package files

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type redisRepository struct {
	conn *redis.Client
}

func NewRedisRepository(conn *redis.Client) Repository {
	return &redisRepository{
		conn: conn,
	}
}

func (r *redisRepository) Get(uuid uuid.UUID) (File, error) {
	return File{}, nil
}

func (r *redisRepository) Set(file File) error {
	return nil
}
