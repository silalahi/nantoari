package main

import (
	"github.com/go-redis/redis"
	"github.com/labstack/echo/v4"
	"github.com/silalahi/nantoari/app/files"
)

func main() {
	r := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	if _, err := r.Ping().Result(); err != nil {
		panic("could not connect to redis")
	}

	e := echo.New()
	e.HideBanner = true
	e.Debug = true

	files.RegisterServiceProvider(e, r)

	e.Logger.Fatal(e.Start(":8080"))
}
