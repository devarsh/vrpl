package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type Address struct {
	gorm.Model
	Line1         string  `gorm:"type:varchar(60)" validate:"required,lte=60" json:"line1"`
	Line2         string  `gorm:"type:varchar(60)" validate:"required,lte=60" json:"line2"`
	Line3         string  `gorm:"type:varchar(60)" validate:"required,lte=60" json:"line3"`
	Line4         string  `gorm:"type:varchar(60)" validate:"required,lte=60" json:"line4"`
	Pincode       string  `gorm:"type:varchar(20)" validate:"required,lte=20" json:"pincode"`
	AreaID        uint    `gorm:"type:int(10) unsigned" validate:"required" json:"area_id"`
	Area          Area    `validate:"-" json:"area"`
	CityID        uint    `gorm:"type:int(10) unsigned" validate:"required" json:"city_id"`
	City          City    `validate:"-" json:"city"`
	StateID       uint    `gorm:"type:int(10) unsigned" validate:"required" json:"state_id"`
	State         State   `validate:"-" json:"state"`
	CountryID     uint    `gorm:"type:int(10) unsigned" validate:"required" json:"country_id"`
	Country       Country `validate:"-" json:"country"`
	AddressTypeID uint    `gorm:"type:int(10) unsigned" validate:"required" json:"address_type_id"`
	AddressType   Types   `validate:"-" json:"address_type"`
	ClientID      *uint   `gorm:"type:int(10) unsigned"  json:"client_id"`
	EmployeeID    *uint   `gorm:"type:int(10) unsigned"  json:"employee_id,omitempty"`
	CompanyID     *uint   `gorm:"type:int(10) unsigned"  json:"company_id,omitempty"`
}

func (u *Address) Bind(r *http.Request) error {
	return nil
}

func AddressStructLevelValidator(sl validator.StructLevel) {
	address := sl.Current().Interface().(Address)

	if address.ClientID == nil && address.EmployeeID == nil && address.CompanyID == nil {
		sl.ReportError(address.ClientID, "client_id", "Address.ClientID", "required", "ClientId or EmployeeId or CompanyId")
		sl.ReportError(address.EmployeeID, "employee_id", "Address.EmployeeID", "required", "ClientId or EmployeeId or CompanyId")
		sl.ReportError(address.CompanyID, "company_id", "Address.CompanyID", "required", "ClientId or EmployeeId or CompanyId")
	}
}
