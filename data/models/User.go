package data

import uuid "github.com/satori/go.uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username  string
	Questions []Question
}

const (
	USERNAME_FIELD_NAME = "username"
	SUPER_USER_NAME     = "admin"
)
