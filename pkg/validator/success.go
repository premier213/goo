package validator

// handle validate request success
func SuccessRequest(message string) *ResponseType {
	var res *ResponseType
	res = &ResponseType{
		Success: true,
		Msg:     message,
	}

	return res
}
