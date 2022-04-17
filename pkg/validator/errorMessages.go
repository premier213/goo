package validator

var errorRequest = map[string]map[string]map[string]string{
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

var errorDatabase = map[string]map[string]string{
	"users_username_key": {
		"23505": "username is Already exist",
		"code":  "1400",
	},
}
