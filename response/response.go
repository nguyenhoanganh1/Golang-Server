package response

type Response struct {
	Status  bool        `json : "status"`
	Message string      `json : "message"`
	Error   string      `json : "error"`
	Data    interface{} `json : "data"`
}

func BuildCommonReponse(status bool, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func BuildErrorResponse(message string, err string) Response {
	return Response{
		Status:  false,
		Message: message,
		Error:   err,
	}
}
