package controllers

// SuccessResponse
// @Description Object that represents a success response
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data"`
	Meta    interface{} `json:"meta,omitempty"`
} // @name SuccessResponse

func NewSuccessResponse(status int, data interface{}) *SuccessResponse {
	return &SuccessResponse{
		Status: status,
		Data:   data,
	}
}
