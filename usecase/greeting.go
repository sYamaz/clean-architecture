package usecase

import "context"

type (
	greetingService struct {
		presenter GreetingPresenter
	}

	GreetingService interface {
		Hello(ctx context.Context, name string) error
	}

	GreetingPresenter interface {
		Reply(ctx context.Context, name string) error
	}
)

func NewGreetingService(presenter GreetingPresenter) GreetingService {
	return &greetingService{
		presenter: presenter,
	}
}

func (s *greetingService) Hello(ctx context.Context, name string) error {
	return s.presenter.Reply(ctx, name)
}
