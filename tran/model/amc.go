package model

import (
	"github.com/devarsh/vrpl/misc/model"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Amc struct {
	gorm.Model
	CompanyID         uint        `gorm:"type:int(10) unsigned" validate:"required" json:"company_id"`
	MachineID         uint        `gorm:"type:int(10) unsigned" validate:"required" json:"machine_id"`
	PaymentIntervalID uint        `gorm:"type:int(10) unsigned" validate:"required" json:"payment_interval_id"`
	PaymentInterval   model.Types `validate:"-" json:"payment_interval"`
	PaymentPeriodID   uint        `gorm:"type:int(10) unsigned" validate:"required" json:"payment_period_id"`
	PaymentPeriod     model.Types `validate:"-" json:"payment_period"`
	AmcTypeID         uint        `gorm:"type:int(10) unsigned" validate:"required" json:"amc_type_id"`
	AmcType           model.Types `validate:"-" json:"amc_type"`
	Amount            float64     `gorm:"type:float(10,2)" validate:"required" json:"amount"`
	FromDt            time.Time   `gorm:"type:timestamp" validate:"ltfield=ToDt" json:"from_dt"`
	ToDt              time.Time   `gorm:"type:timestamp" validate:"gtfield=FromDt" json:"to_dt"`
	ApprovalDt        time.Time   `gorm:"type:timestamp" json:"approval_dt"`
	AmcCancelledDt    *time.Time  `gorm:"type:timestamp" json:"cancelled_dt,omitempty"`
}

func (u *Amc) Bind(r *http.Request) error {
	return nil
}
