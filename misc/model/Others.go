package model

import (
	"net/http"
)

type QueryReq struct {
	Name string `validate:"required,email" json:"username"`
}

func (u *QueryReq) Bind(r *http.Request) error {
	return nil
}

type GetChildrensReq struct {
	ParentID uint `validate:"required" json:"id,string"`
}

func (u *GetChildrensReq) Bind(r *http.Request) error {
	return nil
}
