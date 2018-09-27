package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Complain struct {
	gorm.Model
	MachineID     uint      `gorm:"type:int(10) unsigned" validate:"required" json:"machine_id"`
	ComplainDt    time.Time `gorm:"timestamp" validate:"required" json:"complain_dt"`
	ProblemDesc   string    `gorm:"varchar(100)" validate:"required,lte=100" json:"problem_desc"`
	ContactPerson string    `gorm:"varchar(30)" validate:"required,lte=30" json:"contact_person"`
	ContactNo     string    `gorm:"varchar(30)" validate:"required,lte=30" json:"contact_no"`
}
