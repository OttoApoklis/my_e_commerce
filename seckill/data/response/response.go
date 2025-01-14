package response

type Response struct {
	message string
	data    interface{}
}

func GetResponse(message string, data interface{}) Response {
	response := Response{}
	response.message = message
	response.data = data
	return response
}
