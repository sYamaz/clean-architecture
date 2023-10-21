package presenter

import "context"

type (
	ResultSender interface {
		SendJson(ctx context.Context, i interface{}) error
	}
)
