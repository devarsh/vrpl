package resp

import (
	customErr "github.com/devarsh/vrpl/error"
	"github.com/go-chi/render"
	"net/http"
)

type Resp struct {
	Result interface{} `json:"result"`
	Status bool        `json:"status"`
}

func (u *Resp) Render(w http.ResponseWriter, r *http.Request) error {
	if u.Result != nil {
		if val, ok := u.Result.(*customErr.ErrResponse); ok {
			render.Status(r, val.HTTPStatusCode)
			u.Status = false
			return nil
		} else if _, ok := u.Result.(error); ok {
			u.Status = false
			render.Status(r, http.StatusBadRequest)
		}
		render.Status(r, http.StatusOK)
		u.Status = true
		return nil
	}
	return nil
}
