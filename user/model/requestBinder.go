package model

import (
	"net/http"
)

func (u *User) Bind(r *http.Request) error {
	return nil
}

func (u *UserLogin) Bind(r *http.Request) error {
	return nil
}

func (u *UserPasswordChange) Bind(r *http.Request) error {
	return nil
}

func (u *UserStatus) Bind(r *http.Request) error {
	return nil
}
