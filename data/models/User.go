package data

type User struct {
	Username string `gorm:"primary_key;"`
}

const (
	USERNAME_FIELD_NAME = "username"
	SUPER_USER_NAME     = "admin"
)
