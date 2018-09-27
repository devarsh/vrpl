package error

import (
	"fmt"
	"net/http"
)

func UserPassInvalid() error {
	err := fmt.Errorf("Invalid Username or Password ")
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Kind:           Authentication,
		ErrorText:      "USER_PWD_INVALID",
	}
}

func UserDisable() error {
	err := fmt.Errorf("User is disabled")
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Kind:           Authentication,
		ErrorText:      "USER_DISABLE",
	}
}
