package request

import (
	"net/http"
)

// Response represents a HTTP response.
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

// NewResponse returns a new Response instance.
func NewResponse(status int, message string, data, error any) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   error,
	}
}

// bindJSONErr returns a new Response instance for failed JSON binding.
func bindJSONErr(err error) Response {
	return NewResponse(http.StatusBadRequest, "Failed to bind JSON", nil, err.Error())
}

// bindQueryErr returns a new Response instance for failed query binding.
func bindQueryErr(err error) Response {
	return NewResponse(http.StatusBadRequest, "Failed to bind query", nil, err.Error())
}

// bindPathParamErr returns a new Response instance for failed path param binding.
func bindPathParamErr(err error) Response {
	return NewResponse(http.StatusBadRequest, "Failed to bind path param", nil, err.Error())
}
