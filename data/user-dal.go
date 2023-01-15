package data

import (
	models "questions/data/models"

	"gorm.io/gorm"
)

type UsersDal struct {
	db *gorm.DB
}

func NewUsersDal(db *gorm.DB) UsersDal {
	db.AutoMigrate(&models.User{})

	// Hydrate the super user
	superUser := models.User{
		Username: models.SUPER_USER_NAME,
	}

	db.Create(&superUser)

	return UsersDal{db}
}
