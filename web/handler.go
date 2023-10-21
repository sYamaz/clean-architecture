package web

import (
	"github.com/labstack/echo/v4"
	"github.com/sYamaz/clean-architecture/controller"
)

type (
	handler struct {
		operator ContextOperator
	}

	Handler interface {
		Func(controllerFunc controller.ControllerFunc) echo.HandlerFunc
	}
)

func NewWebHandler(operator ContextOperator) Handler {
	return &handler{
		operator: operator,
	}
}

func (h *handler) Func(controllerFunc controller.ControllerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := h.operator.Wrap(c.Request().Context(), c)
		return controllerFunc(ctx)
	}
}
