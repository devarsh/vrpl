package model

import (
	"github.com/go-chi/render"
	"net/http"
)

type QueryReq struct {
	Name string `validate:"required" json:"name"`
}

type MiscCreatedResp struct {
	ID     uint   `json:"id,omitempty"`
	Status string `json:"status"`
}

type GetChildrensReq struct {
	ParentID uint `validate:"required" json:"id,string"`
}

func (u *QueryReq) Bind(r *http.Request) error {
	return nil
}

func (u *GetChildrensReq) Bind(r *http.Request) error {
	return nil
}

func (u *MiscCreatedResp) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}
