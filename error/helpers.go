package error

import (
	"net/http"

	"github.com/go-chi/render"
)

func ErrInternalError(err error, tag ...string) error {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		Kind:           InternalError,
		ErrorText:      err.Error(),
	}
}

func ErrInvalidRequestError(err error) error {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Kind:           Invalid,
		ErrorText:      err.Error(),
	}
}

func ErrInvalidRequestWithValueError(err error, value interface{}) error {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Kind:           Validation,
		ErrorText:      err.Error(),
		Value:          value,
	}
}

//var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

func UnwrapError(err error) render.Renderer {
	val, ok := err.(*ErrResponse)
	if ok {
		return val
	}
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusInternalServerError,
		Kind:           Unknown,
		ErrorText:      err.Error(),
	}
}

func UnWrapErrorStruct(err error) *ErrResponse {
	val, ok := err.(*ErrResponse)
	if ok {
		return val
	}
	return nil
}
