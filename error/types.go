package error

import (
	"github.com/go-chi/render"
	"net/http"
)

type Kind uint8

const (
	Other Kind = iota
	Validation
	NotFound
	Exists
	Unknown
	Authentication
	InternalError
	Invalid
)

func (k Kind) String() string {
	switch k {
	case Other:
		return "other error"
	case Validation:
		return "validation error"
	case NotFound:
		return "item does not exist"
	case Exists:
		return "item already exists"
	case Unknown:
		return "An unknown error occured"
	case Authentication:
		return "Authentication error occured"
	case InternalError:
		return "An internal Error occured"
	case Invalid:
		return "Invalid operation performed"
	default:
		return "This is not a valid kind"
	}
}

type ErrResponse struct {
	Err            error       `json:"-"`
	HTTPStatusCode int         `json:"-"`
	Kind           Kind        `json:"kind,omitempty"`
	ErrorText      string      `json:"error,omitempty"`
	Value          interface{} `json:"value,omitempty"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func (e *ErrResponse) Error() string {
	return e.Err.Error()
}
