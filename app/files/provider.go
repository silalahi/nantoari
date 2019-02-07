package files

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
)

// RegisterServiceProvider register file service handler
func RegisterServiceProvider(e *echo.Echo, redis *redis.Client) {
	NewHTTPHandler(e,
		NewUsecase(
			NewRedisRepository(redis),
		),
	)
}
