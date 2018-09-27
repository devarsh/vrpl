package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

//This model helps define various miscallenious into single table, seperated by groups

type Types struct {
	gorm.Model
	GroupNm   string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"-"`
	Name      string `gorm:"type:varchar(30)" validate:"required,lte=30" json:"name"`
	ShortName string `gorm:"type:varchar(10)" json:"short_name,omitempty"`
}

func (u *Types) Bind(r *http.Request) error {
	return nil
}
