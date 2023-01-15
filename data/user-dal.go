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

func (dal *UsersDal) GetByUsername(username string) (models.User, error) {
	user := models.User{}

	if err := dal.db.Where("username = ?", username).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
