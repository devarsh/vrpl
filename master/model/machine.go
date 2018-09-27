package model

import (
	"github.com/devarsh/vrpl/misc/model"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Machine struct {
	gorm.Model
	ClientID   uint       `gorm:"type:int(10) unsigned" validate:"required" json:"client_id"`
	ItemID     uint       `gorm:"type:int(10) unsigned" validate:"required" json:"item_id"`
	Item       model.Item `validate:"-" json:"item"`
	SerialNo   string     `gorm:"type:varchar(20)" validate:"lte=20" json:"serial_no"`
	Remarks    string     `gorm:"type:varchar(100)" validate:"lte=100" json:"remarks"`
	EmployeeID uint       `gorm:"type:int(10) unsigned" validate:"required" json:"employee_id"`
	Cancelled  *time.Time `gorm:"type:timestamp" json:"cancelled,omitempty"`
}

func (u *Machine) Bind(r *http.Request) error {
	return nil
}
