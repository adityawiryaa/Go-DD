package valueobject

import (
	"errors"
	"fmt"
)

type Name struct {
	first string
	last  string
}

func (n Name) First() string {
	return n.first
}

func (n Name) Last() string {
	return n.last
}

func (n Name) Full() string {
	if n.last == "" {
		return n.First()
	}
	return fmt.Sprintf("%s %s", n.first, n.last)
}

func (n Name) IsEmpty() bool {
	return n.First() == ""
}

func NewName(first, last string) (Name, error) {
	if first == "" {
		return Name{}, errors.New("first name cannot be empty")
	}
	return Name{first, last}, nil
}
