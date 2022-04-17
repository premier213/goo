package validator

import (
	"github.com/go-playground/validator/v10"
)

type ResponseType struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Code    string `json:"code,omitempty"`
}

var isValid = validator.New()

func ValidateRequest(v interface{}) *ResponseType {
	var res *ResponseType
	err := isValid.Struct(v)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			res = ErrorRequest(err.Field(), err.Tag())
		}
	}

	return res
}
