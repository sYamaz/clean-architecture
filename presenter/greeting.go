package presenter

import (
	"context"
	"fmt"

	"github.com/sYamaz/clean-architecture/usecase"
)

type (
	greetingPresenter struct {
		sender ResultSender
	}

	GreetingPresenter interface {
		Reply(ctx context.Context, name string) error
	}
)

func NewGreetingPresenter(sender ResultSender) usecase.GreetingPresenter {
	return &greetingPresenter{
		sender: sender,
	}
}

func (p *greetingPresenter) Reply(ctx context.Context, name string) error {
	type Response struct {
		Greeting string `json:"greeting"`
	}
	res := Response{
		Greeting: fmt.Sprintf("hello %v", name),
	}

	return p.sender.SendJson(ctx, &res)
}
