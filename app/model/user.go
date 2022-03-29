package model

import hook "goo/platform/hooks"

type User struct {
	hook.Base
	Username string `json:"username" validate:"required,min=3"    gorm:"unique;not null;"`
	Password string `json:"password"`
}
