package main

import (
	"github.com/go-redis/redis"
	"github.com/silalahi/nantoari/app/files"

	"github.com/labstack/echo/v4"
)

func main() {
	r := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	e := echo.New()
	e.HideBanner = true
	e.Debug = true

	files.RegisterServiceProvider(e, r)

	e.Logger.Fatal(e.Start(":8080"))
}
