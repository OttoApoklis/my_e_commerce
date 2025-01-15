package response

type Response struct {
	Code    int
	Message string
	Data    interface{}
}

func GetResponse(code int, message string, data interface{}) Response {
	response := Response{}
	response.Code = code
	response.Message = message
	response.Data = data
	return response
}
