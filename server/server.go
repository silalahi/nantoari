package server

import (
	"context"
	"github.com/silalahi/nantoari/file"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/silalahi/nantoari/store"
)

type Server struct {
	http *echo.Echo
	config *Config
	db *store.KV
}

func New(cfg *Config) (*Server, error) {
	server := &Server{
		http: echo.New(),
		config: cfg,
	}

	err := server.Init()
	if err != nil {
		return nil, err
	}

	return server, nil
}

func (s *Server) UseStore(db store.KV) {
	s.db = &db
}

func (s *Server) Init() error {
	s.http.HideBanner = true
	s.http.Debug = true

	s.http.Use(middleware.Recover())
	s.http.Use(middleware.Logger())

	s.http.GET("/:uuid", file.HTTPGetHandler)
	s.http.POST("/", file.HTTPSetHandler)

	return nil
}

func (s *Server) Run() error {
	go func() {
		if err := s.http.Start(s.config.Addr()); err != nil {
			s.http.Logger.Info("shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		s.http.Logger.Fatal(err)
	}

	<-ctx.Done()

	return nil
}