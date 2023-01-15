package data

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username  string    `gorm:"unique;not null;"`
	Questions []Question
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.ID = uuid.NewV4()

	return nil
}

const (
	USER_HEADER_ID  = "id"
	SUPER_USER_NAME = "admin"
)
