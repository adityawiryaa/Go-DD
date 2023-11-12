package greeting_usecase

import "go-driven-design/domain/usecase"

type GreetingService struct{}

func NewGreetingService() usecase.GreetingUseCase {
	return &GreetingService{}
}
