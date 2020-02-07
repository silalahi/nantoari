package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	*echo.Echo
}

func NewHTTPServer(cfg *Config) (*HTTPServer, error) {
	server := &HTTPServer{
		Echo: echo.New(),
	}

	server.Echo.Use(middleware.Logger())
	server.Echo.Use(middleware.Recover())

	return server, nil
}
