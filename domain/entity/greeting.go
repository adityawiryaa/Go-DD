package entity

import "go-driven-design/domain/valueobject"

type GreetingDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Greeting struct {
	Name valueobject.Name
}

func NewGreeting(payload *GreetingDTO) (*Greeting, error) {
	name, err := valueobject.NewName(payload.FirstName, payload.LastName)
	if err != nil {
		return nil, err
	}
	return &Greeting{
		Name: name,
	}, nil
}
