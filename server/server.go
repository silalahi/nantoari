package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	http *echo.Echo
	config *Config
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

func (s *Server) Init() error {
	s.http.HideBanner = true
	s.http.Debug = true

	s.http.Use(middleware.Recover())
	s.http.Use(middleware.Logger())

	// TODO: enable pprof http. currently we must create http handler wrapper

	s.http.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	for _, r := range pprofRoutes {
		switch r.Method {
		case "GET":
			s.http.GET(r.Path, r.Handler)
		case "POST":
			s.http.POST(r.Path, r.Handler)
		}
	}

	return nil
}

func (s *Server) Run() error {
	s.http.Logger.Fatal(s.http.Start(s.config.Addr()))

	return nil
}

// func (s *Server) Run() error {
// 	go func() {
// 		if err := srv.Start(":" + os.Getenv("APP_PORT")); err != nil {
// 			srv.Logger.Info("shutting down the server")
// 		}
// 	}()

// 	quit := make(chan os.Signal)
// 	signal.Notify(quit, os.Interrupt)
// 	<-quit

// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()

// 	if err := srv.Shutdown(ctx); err != nil {
// 		srv.Logger.Fatal(err)
// 	}

// 	<-ctx.Done()
// }
