package domain

import "fmt"

type UserSnapshot struct {
	id    string
	name  string
	email string
}

func NewUserSnapshot(id, name, email string) (*UserSnapshot, error) {
	if id == "" {
		return nil, fmt.Errorf("id cannot be empty")
	}
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if email == "" {
		return nil, fmt.Errorf("email cannot be empty")
	}
	return &UserSnapshot{
		id:    id,
		name:  name,
		email: email,
	}, nil
}
