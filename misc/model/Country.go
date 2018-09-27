package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Country struct {
	gorm.Model
	Name string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
}

func (u *Country) Bind(r *http.Request) error {
	return nil
}
