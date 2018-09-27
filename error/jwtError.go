package error

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)
import "net/http"

func JwtTokenBlackList() error {
	err := fmt.Errorf("Token Blacklisted")
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Kind:           Authentication,
		ErrorText:      "TOKEN_EXPIRED",
	}
}

func JwtError(err error) error {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return &ErrResponse{
				Err:            err,
				HTTPStatusCode: http.StatusBadRequest,
				Kind:           Authentication,
				ErrorText:      "NOT_VALID_TOKEN",
			}
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return &ErrResponse{
				Err:            err,
				HTTPStatusCode: http.StatusBadRequest,
				Kind:           Authentication,
				ErrorText:      "TOKEN_EXPIRED",
			}
		}
	}
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Kind:           Authentication,
		ErrorText:      "TOKEN_UNKNOWN_ERR",
	}
}
