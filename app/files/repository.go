package files

import (
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

// redisRepository is implementation of Repository using Redis
type redisRepository struct {
	conn *redis.Client
}

// NewRedisRepository returns an instance of Repository implementation
// using Redis
func NewRedisRepository(conn *redis.Client) Repository {
	return &redisRepository{
		conn: conn,
	}
}

// Get retrieves File from repository by given UUID.
func (r *redisRepository) Get(id uuid.UUID) (File, error) {
	url, err := r.conn.Get(id.String()).Result()
	if err != nil {
		return File{}, ErrNotFound
	}

	return File{
		uuid: id,
		url:  url,
	}, nil
}

// Set stores File to repository.
func (r *redisRepository) Set(file File) error {
	return r.conn.Set(file.UUID().String(), file.URL(), 0).Err()
}
