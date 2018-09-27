package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Area struct {
	gorm.Model
	Name   string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
	CityID uint   `gorm:"type:int(10) unsigned" validate:"required" json:"city_id"`
}

func (u *Area) Bind(r *http.Request) error {
	return nil
}
