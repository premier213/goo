package models

import hook "goo/pkg/hooks"

type User struct {
	hook.Base
	Username string `json:"username"`
	Password string `json:"password"`
}
