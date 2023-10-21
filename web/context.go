package web

import (
	"context"

	"github.com/labstack/echo/v4"
)

type (
	echoContextOperator struct{}

	ContextOperator interface {
		Wrap(ctx context.Context, c echo.Context) context.Context
		Unwrap(ctx context.Context) echo.Context
	}

	ContextKey struct{}
)

func NewContextOperator() ContextOperator {
	return &echoContextOperator{}
}

func (o *echoContextOperator) Wrap(ctx context.Context, c echo.Context) context.Context {
	return context.WithValue(ctx, ContextKey{}, c)
}

func (o *echoContextOperator) Unwrap(ctx context.Context) echo.Context {
	return ctx.Value(ContextKey{}).(echo.Context)
}
