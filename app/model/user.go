package model

import hook "goo/platform/hooks"

type Mp4 struct {
	Username string `json:"username" query:"username" validate:"required,min=3" gorm:"unique;not null"`
	Password string `json:"password"`
}

type User struct {
	hook.Base
	Mp4
}
