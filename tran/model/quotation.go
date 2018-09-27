package model

import (
	"github.com/devarsh/vrpl/misc/model"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type Quotation struct {
	gorm.Model
	CompanyID         uint        `gorm:"type:int(10) unsigned" validate:"required" json:"company_id"`
	ClientID          uint        `gorm:"type:int(10) unsigned" validate:"required" json:"client_id"`
	QuotationDt       time.Time   `gorm:"timestamp" validate:"required" json:"quotation_dt"`
	PaymentIntervalID uint        `gorm:"type:int(10) unsigned" validate:"required" json:"payment_interval_id"`
	PaymentInterval   model.Types `validate:"-" json:"payment_interval"`
	PaymentPeriodID   uint        `gorm:"type:int(10) unsigned" validate:"required" json:"payment_period_id"`
	PaymentPeriod     model.Types `validate:"-" json:"payment_period"`
	AmcTypeID         uint        `gorm:"type:int(10) unsigned" validate:"required" json:"amc_type_id"`
	AmcType           model.Types `validate:"-" json:"amc_type"`
	FromDt            time.Time   `gorm:"type:timestamp" validate:"ltfield=ToDt" json:"from_dt"`
	ToDt              time.Time   `gorm:"type:timestamp" validate:"gtfield=FromDt" json:"to_dt"`
	ApprovalDt        time.Time   `gorm:"type:timestamp" json:"approval_dt"`
}

func (u *Quotation) Bind(r *http.Request) error {
	return nil
}

type QuotationDtl struct {
	gorm.Model
	QuotationID uint    `gorm:"type:int(10) unsigned" validate:"required" json:"company_id"`
	MachineID   uint    `gorm:"type:int(10) unsigned" validate:"required" json:"machine_id"`
	Amount      float64 `gorm:"type:float(10,2)" validate:"required" json:"amount"`
	AmcID       uint    `gorm:"type:int(10) unsigned" validate:"required" json:"amc_id"`
}

func (u *QuotationDtl) Bind(r *http.Request) error {
	return nil
}
