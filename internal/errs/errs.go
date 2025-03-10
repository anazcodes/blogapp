// Package errs implements error mapping utilities.
package errs

import (
	"errors"

	"github.com/anazcodes/blogapp/pkg/request"
	"github.com/gofiber/fiber/v2"
)

// InternalError creates a request.Response for internal server errors.
// It returns a response with a 500 status code and a generic error message.
func InternalError(err error) request.Response {
	return request.Response{
		Status:  fiber.StatusInternalServerError,
		Message: "Failed to process your request",
		Error:   err.Error(),
	}
}

// UnwrapAll recursively unwraps an error until it reaches the root cause.
// It returns the innermost (root) error.
func UnwrapAll(err error) error {
	for {
		unwrapped := errors.Unwrap(err)
		if unwrapped == nil {
			return err
		}
		err = unwrapped
	}
}
