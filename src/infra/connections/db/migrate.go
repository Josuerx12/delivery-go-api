package db

import (
	"delivery-go-api/src/domain/entities"

	"gorm.io/gorm"
)

func Migrate(database *gorm.DB) error {
	err := database.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}

	return nil
}