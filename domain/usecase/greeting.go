package usecase

import (
	"context"
	"go-driven-design/domain/entity"
)

type GreetingUseCase interface {
	Greet(ctx context.Context, payload *entity.GreetingDTO) (string, error)
}
