package main

import (
	"go-driven-design/application" // replace with your module path
	"go-driven-design/domain/usecase"
	"go-driven-design/infrastructure"
	greeting_usecase "go-driven-design/internal/usecase/greeting"
)

type UseCaseType struct {
	GreetingUseCase usecase.GreetingUseCase
}

func main() {
	uc := initUseCase()
	greetHandler := application.NewGreetHandler(uc.GreetingUseCase)
	infrastructure.StartServer(greetHandler)
}

func initUseCase() *UseCaseType {
	greetingService := greeting_usecase.NewGreetingService()
	return &UseCaseType{
		GreetingUseCase: greetingService,
	}
}
