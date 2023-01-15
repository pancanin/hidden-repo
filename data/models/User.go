package data

import uuid "github.com/satori/go.uuid"

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username  string
	Questions []Question
}

const (
	USER_HEADER_ID  = "id"
	SUPER_USER_NAME = "admin"
)
