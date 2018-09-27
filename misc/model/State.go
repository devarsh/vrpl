package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type State struct {
	gorm.Model
	Name      string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
	CountryID uint   `gorm:"type:int(10) unsigned" validate:"required" json:"country_id"`
}

func (u *State) Bind(r *http.Request) error {
	return nil
}
