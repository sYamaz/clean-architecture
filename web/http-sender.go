package web

import (
	"context"
	"net/http"

	"github.com/sYamaz/clean-architecture/presenter"
)

type (
	httpResultSender struct {
		operator ContextOperator
	}
)

func NewResultSender(operator ContextOperator) presenter.ResultSender {
	return &httpResultSender{
		operator: operator,
	}
}

func (s *httpResultSender) SendJson(ctx context.Context, i interface{}) error {
	c := s.operator.Unwrap(ctx)
	return c.JSON(http.StatusOK, i)
}
