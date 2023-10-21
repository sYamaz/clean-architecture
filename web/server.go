package web

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type (
	server struct {
		e *echo.Echo
	}

	Server interface {
		Run()
	}
)

func NewServer(router Router) Server {
	return &server{
		e: router.Echo(),
	}
}

func (s *server) Run() {
	s.e.Use(middleware.RequestID())

	s.e.Logger.Fatal(s.e.Start(":3000"))
}
