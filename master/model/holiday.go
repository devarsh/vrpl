package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Holiday struct {
	gorm.Model
	Name string    `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
	Date time.Time `gorm:"type:timestamp" validate:"required" json:"date"`
}

func (u *Holiday) Bind(r *http.Request) error {
	return nil
}
