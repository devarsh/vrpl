package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Attendance struct {
	gorm.Model
	EmployeeID uint      `gorm:"type:int(10) unsigned" validate:"required" json:"employee_id"`
	Date       time.Time `gorm:"type:timestamp" validate:"required" json:"date"`
	Status     uint8     `gorm:"type:tinyint(1)" validate:"required" json:"name"`
	Reason     string    `gorm:"type:varchar(50)" validate:"lte=50" json:"reason"`
}

func (u *Attendance) Bind(r *http.Request) error {
	return nil
}

type AttendanceReq struct {
	gorm.Model
	EmployeeID uint      `gorm:"type:int(10) unsigned" validate:"required" json:"employee_id"`
	FromDt     time.Time `gorm:"type:timestamp" validate:"ltfield=ToDt" json:"from_dt"`
	ToDt       time.Time `gorm:"type:timestamp" validate:"gtfield=FromDt" json:"to_dt"`
	Reason     string    `gorm:"type:varchar(50)" validate:"required,lte=50" json:"reason"`
}

func (u *AttendanceReq) Bind(r *http.Request) error {
	return nil
}
