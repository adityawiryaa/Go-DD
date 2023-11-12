package greeting_usecase

import (
	"context"
	"fmt"
	"go-driven-design/domain/entity"
	"go-driven-design/pkg/common"
	"log"
)

func (s *GreetingService) Greet(ctx context.Context, payload *entity.GreetingDTO) (string, error) {
	result, err := entity.NewGreeting(payload)
	if err != nil {
		log.Println(common.DebugType{
			Property: "[ERROR] NEW GREETING ENTITY",
			Error:    err,
			Function: "NewGreeting",
		})
		return "", err
	}
	return fmt.Sprintf("Hello, %s", result.Name.Full()), nil
}
