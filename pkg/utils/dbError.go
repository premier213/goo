package utils

type Mp map[string]interface{}

var d = map[string]map[string]string{
	"users_username_key": {
		"23505": "username is Already exist",
		"code":  "1400",
	},
}

var n = map[string]map[string]string{
	"email_not_valid": {
		"msg":  "Email Not Valid",
		"code": "1402",
	},
	"not_send": {
		"msg":  "could not send email",
		"code": "1403",
	},
}

/**
 * @params c database code error
 * @params f database constraint name error
 * return database error message with code and constraint name error
 */
func DbError(c, f string) *Mp {
	res := Mp{
		"Error": true,
		"Msg":   d[f][c],
		"Code":  d[f]["code"],
	}

	return &res
}
func EmailError(f string) *Mp {
	res := Mp{
		"Error": true,
		"Msg":   n[f]["msg"],
		"Code":  n[f]["code"],
	}

	return &res
}
