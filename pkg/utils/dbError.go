package utils

type Mp map[string]interface{}

var d = map[string]map[string]string{
	"users_username_key": {
		"23505": "username is Already exist",
		"code":  "1400",
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
