package controller

import "context"

type (
	ControllerFunc func(ctx context.Context) error
)
