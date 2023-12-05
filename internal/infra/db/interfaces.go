package db

import "github.com/RomeroGabriel/go-api-standards/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
