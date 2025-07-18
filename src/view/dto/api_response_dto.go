package dto

type APIResponse struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewAPIResponse(statusCode int, message string, data interface{}) APIResponse {
	return APIResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
