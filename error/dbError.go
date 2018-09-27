package error

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

func DBRecordNotFound() error {
	err := fmt.Errorf("Record not found in the database")
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusNotFound,
		Kind:           NotFound,
		ErrorText:      NotFound.String(),
	}
}

func ErrMySqlDBError(err error) error {
	if gorm.IsRecordNotFoundError(err) {
		return &ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusNotFound,
			Kind:           NotFound,
			ErrorText:      NotFound.String(),
		}
	} else if val, ok := err.(*mysql.MySQLError); ok {
		switch val.Number {
		case 1062:
			return &ErrResponse{
				Err:            err,
				HTTPStatusCode: http.StatusBadRequest,
				Kind:           Exists,
				ErrorText:      Exists.String(),
			}
		}
	}
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: http.StatusBadRequest,
		Kind:           Unknown,
		ErrorText:      err.Error(),
	}
}
