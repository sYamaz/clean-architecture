package customerror

import "fmt"

type (
	BindError struct {
		Internal error
	}
)

func NewBindError(err error) error {
	return &BindError{
		Internal: err,
	}
}

func (e *BindError) Error() string {
	return fmt.Sprintf("bind error: %v", e.Internal)
}
