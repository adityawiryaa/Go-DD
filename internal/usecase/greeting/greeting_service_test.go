package greeting_usecase_test

import (
	"context"
	"errors"
	"fmt"
	"go-driven-design/domain/entity"
	greeting_usecase "go-driven-design/internal/usecase/greeting"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreet(t *testing.T) {
	type args struct {
		ctx   context.Context
		input *entity.GreetingDTO
	}
	service := greeting_usecase.NewGreetingService()
	greeting, _ := entity.NewGreeting(&entity.GreetingDTO{
		FirstName: "Andi",
		LastName:  "Surya",
	})
	tests := []struct {
		name       string
		args       args
		beforeTest func()
		afterTest  func()
		expected   string
		errors     error
	}{
		{
			name: "should be success create admin approval with type create user",
			args: args{
				ctx: context.TODO(),
				input: &entity.GreetingDTO{
					FirstName: "Andi",
					LastName:  "Surya",
				},
			},
			expected: fmt.Sprintf("Hello, %s", greeting.Name.Full()),
		},
		{
			name: "should be failed first name is required",
			args: args{
				ctx: context.TODO(),
				input: &entity.GreetingDTO{
					FirstName: "",
					LastName:  "Surya",
				},
			},
			errors: errors.New("first name cannot be empty"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := service.Greet(tt.args.ctx, tt.args.input)
			if err != nil {
				if tt.errors == nil {
					t.Errorf("GreetUseCase receive error = %v", err)
				} else {
					if !assert.EqualError(t, err, tt.errors.Error()) {
						t.Errorf("GreetUseCase receive error = %v, expected errors = %v", err, tt.errors)
					}
				}
				if tt.afterTest != nil {
					tt.afterTest()
				}
				return
			}
			if !assert.Equal(t, got, tt.expected) {
				t.Errorf("GreetUseCase receive = %v, expected = %v", got, tt.expected)
			}
			if tt.afterTest != nil {
				tt.afterTest()
			}
		})
	}
}
