package validator

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"goo/pkg/errs"
)

var m = map[string]map[string]map[string]string{
	"Username": {
		"required": {
			"msg":  "username is required",
			"code": "1402",
		},
		"min": {
			"msg":  "username must be at least 3 characters",
			"code": "1403",
		},
	},
}

var validate = validator.New()

/**
 * @params f validation field
 * @params t validation tag
 * return validation struct message with code and constraint name error
 */
func valid(f string, t string) *errs.Mp {
	fmt.Println(f, t)
	res := errs.Mp{
		"Error": true,
		"Msg":   m[f][t]["msg"],
		"Code":  m[f][t]["code"],
	}

	return &res
}

func ValidateStruct(v interface{}) *errs.Mp {
	var res *errs.Mp
	err := validate.Struct(v)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			res = valid(err.Field(), err.Tag())
		}
	}
	return res
}
