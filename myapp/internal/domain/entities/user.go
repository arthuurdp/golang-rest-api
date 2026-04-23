package entities

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID
	Name string
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(name, email, hashedPassword string) (*User, error) {
	if name == "" {
		return nil, errors.New("Name can't be empty")
	}
	
	if email == "" {
		return nil, errors.New("Email can't be empty")
	}

	now := time.Now()

	return &User {
		ID: uuid.New(),
		Name: name,
		Email: email,
		Password: hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}