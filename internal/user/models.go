package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
