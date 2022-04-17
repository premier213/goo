package model

import hook "goo/platform/hooks"

type Mp4 struct {
	Username string `json:"username" query:"username" validate:"required,min=3" gorm:"unique;not null"`
}

type User struct {
	hook.Base
	Mp4
}

type Email struct {
	Email string `json:"email" query:"email" validate:"required,email"`
}
