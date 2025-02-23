package blogapperr

import (
	"net/http"

	"github.com/anazcodes/blog-crud-api/internal/errs"
	"github.com/anazcodes/blog-crud-api/internal/repository/blogrepo/cache"
	"github.com/anazcodes/blog-crud-api/pkg/request"
)

var responses = map[error]request.Response{
	cache.ErrCacheInMaxCap: {
		Status:  http.StatusUnprocessableEntity,
		Error:   cache.ErrCacheInMaxCap.Error(),
		Message: "Failed to save, blog storage capacity reached",
	},
	cache.ErrItemNotFound: {
		Status:  http.StatusNotFound,
		Error:   cache.ErrItemNotFound.Error(),
		Message: "Referenced resource does not found in the system",
	},
}

// Response maps an error to a request.Response.
// It unwraps the error and returns the corresponding response.
// If no specific response is found, it returns an internal server error.
func Response(err error) request.Response {
	uerr := errs.UnwrapAll(err)
	res := corresponding(uerr)

	res.Error = err.Error()

	return res
}

// corresponding returns the request.Response associated with the given error.
// If no matching response is found, it returns an internal server error.
func corresponding(err error) request.Response {
	res, ok := responses[err]
	if !ok {
		return errs.InternalError(err)
	}

	return res
}
