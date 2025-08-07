package controllers

import "net/http"

type ErrorResponse struct {
	Status  int         `json:"status"`
	Error   string      `json:"error"`
	Details interface{} `json:"details,omitempty"`
} // @name ErrorResponse

// NewErrorResponse creates a new Error Response
func NewErrorResponse(status int, details interface{}) ErrorResponse {
	return ErrorResponse{
		Status:  status,
		Error:   http.StatusText(status),
		Details: details,
	}
}
