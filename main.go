package main

import (
	"github.com/sYamaz/clean-architecture/controller"
	"github.com/sYamaz/clean-architecture/presenter"
	"github.com/sYamaz/clean-architecture/usecase"
	"github.com/sYamaz/clean-architecture/web"
)

func main() {
	operator := web.NewContextOperator()
	binder := web.NewBinder(operator)
	sender := web.NewResultSender(operator)

	// presenter
	greetingPresenter := presenter.NewGreetingPresenter(sender)
	// usecase
	greetingService := usecase.NewGreetingService(greetingPresenter)
	// controller
	helloController := controller.NewHelloController(binder, greetingService)

	e := web.NewEcho()
	webHandler := web.NewWebHandler(operator)
	router := web.NewRouter(e, webHandler, helloController)
	server := web.NewServer(router)

	server.Run()
}
