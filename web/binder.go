package web

import (
	"context"

	"github.com/sYamaz/clean-architecture/controller"
	customerror "github.com/sYamaz/clean-architecture/custom-error"
)

type (
	binder struct {
		operator ContextOperator
	}
)

func NewBinder(operator ContextOperator) controller.Binder {
	return &binder{
		operator: operator,
	}
}

func (b *binder) Bind(ctx context.Context, i interface{}) error {
	c := b.operator.Unwrap(ctx)
	// echoのbind機能を使う場合
	if err := c.Bind(i); err != nil {
		return customerror.NewBindError(err)
	}
	return nil
}

// func (b *binder) Bind(ctx context.Context, i interface{}) error {
// 	// 自力で`header:"KEY"`を文字列型としてbindしたい場合の例
// 	v := reflect.ValueOf(i).Elem()
// 	t := reflect.TypeOf(i).Elem()
// 	for index := 0; index < v.NumField(); index++ {
// 		headerKey := t.Field(index).Tag.Get("header")
// 		headerValue := c.Request().Header.Get(headerKey)
// 		v.Field(index).SetString(headerValue)
// 	}
// 	return nil
// }
