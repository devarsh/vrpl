package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Service struct {
	gorm.Model
	Date       time.Time `gorm:"type:timestamp" validate:"ltfield=ToDt" json:"date"`
	ReportNo   uint      `gorm:"type:int(10) unsigned" validate:"required" json:"employee_id"`
	EmployeeID uint      `gorm:"type:int(10) unsigned" validate:"required" json:"employee_id"`
	MachineID  uint      `gorm:"type:int(10) unsigned" validate:"required" json:"machine_id"`
	Remarks    string    `gorm:"type:varchar(50)"  json:"remarks"`
}

func (u *Service) Bind(r *http.Request) error {
	return nil
}
