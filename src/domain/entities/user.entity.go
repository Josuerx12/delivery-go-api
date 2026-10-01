package entities

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" json:"id" validate:"required,uuid"`
	Name      string    `gorm:"size:255;not null" json:"name" validate:"required"`
	Email     string    `gorm:"size:255;unique;not null" json:"email" validate:"required,email"`
	Password  string    `gorm:"size:255;not null" json:"-" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}