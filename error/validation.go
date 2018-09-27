package error

import (
	miscModel "github.com/devarsh/vrpl/misc/model"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	"reflect"
	"strings"
	"time"
)

type ValidationErrorArray map[string][]*ValidationErrorDetail

type ValidationErrorDetail struct {
	Tag      string `json:"tag"`
	Param    string `json:"param,omitempty"`
	ErrorMsg string `json:"errorMsg"`
}

var Validate *validator.Validate

func init() {
	Validate = validator.New()
	Validate.RegisterStructValidation(miscModel.AddressStructLevelValidator, miscModel.Address{})
	Validate.RegisterStructValidation(miscModel.ContactStructLevelValidator, miscModel.Contact{})
	Validate.RegisterTagNameFunc(TagNameFunc)
}

func ValidateStruct(value interface{}) error {
	err := Validate.Struct(value)
	if err != nil {
		customErr := validatorErrorCheck(err)
		return &ErrResponse{
			Err:            err,
			HTTPStatusCode: http.StatusBadRequest,
			Kind:           Validation,
			ErrorText:      Validation.String(),
			Value:          customErr,
		}
	}
	return nil
}

func TagNameFunc(fld reflect.StructField) string {
	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	if name == "-" {
		return ""
	} else if name == "" {
		return fld.Name
	}
	return name
}

func validatorErrorCheck(err error) *ValidationErrorArray {
	errors := make(ValidationErrorArray)
	for _, err := range err.(validator.ValidationErrors) {
		val, ok := err.Value().(string)
		if !ok {
			val1, ok := err.Value().(time.Time)
			if ok {
				val = val1.String()
			} else {
				val = "<<NO!STRING>>"
			}
		}
		var errMsg string
		switch err.Tag() {
		case "gt":
			errMsg = err.Field() + ": Should be greater than " + err.Param() + ", got value=" + val
		case "gte":
			errMsg = err.Field() + ": Should be greater than or equal to " + err.Param() + ", got value=" + val
		case "lt":
			errMsg = err.Field() + ": Should be less than " + err.Param() + ", got value=" + val
		case "lte":
			errMsg = err.Field() + ": Should be less than or equal to " + err.Param() + ", got value=" + val
		case "email":
			errMsg = err.Field() + ": is not a valid email: " + ", got value: " + val
		case "required":
			errMsg = err.Field() + ": is a required field and cannot be empty"
		default:
			errMsg = err.Field() + ": failed validation:" + err.Tag() + "=" + err.Param() + ", got value=" + val
		}
		errDtl := &ValidationErrorDetail{
			Tag:      err.Tag(),
			Param:    err.Param(),
			ErrorMsg: errMsg,
		}
		errors[err.Field()] = append(errors[err.Field()], errDtl)
	}
	return &errors
}
