package controller

import (
	"context"

	"github.com/sYamaz/clean-architecture/usecase"
)

type (
	helloController struct {
		binder  Binder
		service usecase.GreetingService
	}

	HelloController interface {
		Get(ctx context.Context) error
	}
)

func NewHelloController(binder Binder, service usecase.GreetingService) HelloController {
	return &helloController{
		binder:  binder,
		service: service,
	}
}

func (c *helloController) Get(ctx context.Context) error {
	type Request struct {
		Name string `query:"name"`
	}
	req := new(Request)

	if err := c.binder.Bind(ctx, req); err != nil {
		return err
	}

	return c.service.Hello(ctx, req.Name)
}
