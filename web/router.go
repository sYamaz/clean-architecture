package web

import (
	"github.com/labstack/echo/v4"
	"github.com/sYamaz/clean-architecture/controller"
)

type (
	router struct {
		e *echo.Echo
	}

	Router interface {
		Echo() *echo.Echo
	}
)

func NewRouter(e *echo.Echo,
	// webhandler
	handler Handler,
	// controller
	hello controller.HelloController,
	// ...
) Router {
	e.GET("/hello", handler.Func(hello.Get))
	return &router{
		e: e,
	}
}

func (r *router) Echo() *echo.Echo {
	return r.e
}
