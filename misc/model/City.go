package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type City struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
	StateID uint   `gorm:"type:int(10) unsigned" validate:"required" json:"state_id"`
}

func (u *City) Bind(r *http.Request) error {
	return nil
}
