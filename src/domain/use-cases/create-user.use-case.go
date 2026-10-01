package usecases

import (
	"delivery-go-api/src/domain/entities"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name string, email string, password string) (*entities.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entities.User{
		ID:       uuid.New(),
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
	}

	validator := validator.New()

	if err := validator.Struct(user); err != nil {
		return nil, err
	}

	return user, nil
}