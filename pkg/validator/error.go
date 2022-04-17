package validator

// handle validate request error
func ErrorRequest(field, tag string) *ResponseType {
	var res *ResponseType
	res = &ResponseType{
		Success: false,
		Msg:     errorRequest[field][tag]["msg"],
		Code:    errorRequest[field][tag]["code"],
	}

	return res
}

// handle database error
func ErrorDb(code, constraint string) *ResponseType {
	var res *ResponseType
	res = &ResponseType{
		Success: false,
		Msg:     errorDatabase[constraint][code],
		Code:    errorDatabase[constraint]["code"],
	}

	return res
}
