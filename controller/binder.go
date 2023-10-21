package controller

import "context"

type (
	Binder interface {
		Bind(ctx context.Context, i interface{}) error
	}
)
