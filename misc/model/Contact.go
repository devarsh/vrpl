package model

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)


type Contact struct {
	gorm.Model
	ContactTypeID *uint  `gorm:"type:int(10) unsigned" json:"contact_type_id"`
	ContactNo     string `gorm:"type:varchar(20)" json:"contact_no"`
	ContactPerson string `gorm:"type:varchar(50)" json:"contact_person"`
	Designation   string `gorm:"type:varchar(50)" json:"designation"`
	EmailTypeID   *uint  `gorm:"type:int(10) unsigned" json:"email_type_id"`
	Email         string `gorm:"type:varchar(50)" validate:"email" json:"email"`
	SmsEnabled    uint8  `gorm:"type:tinyint(1)" json:"sms_enabled"`
	ContactType   Types  `validate:"-" json:"contact_type"`
	EmailType     Types  `validate:"-" json:"email_type"`
	ClientID      *uint  `gorm:"type:int(10) unsigned"  json:"client_id,omitempty"`
	EmployeeID    *uint  `gorm:"type:int(10) unsigned"  json:"employee_id,omitempty"`
	CompanyID     *uint  `gorm:"type:int(10) unsigned"  json:"company_id,omitempty"`
}

func (u *Contact) Bind(r *http.Request) error {
	return nil
}

func ContactStructLevelValidator(sl validator.StructLevel) {
	contact := sl.Current().Interface().(Contact)
	param := "client_id or employee_id or company_id"
	if contact.ClientID == nil && contact.EmployeeID == nil && contact.CompanyID == nil {
		sl.ReportError(contact.ClientID, "client_id", "Contact.ClientID", "required", param)
		sl.ReportError(contact.EmployeeID, "employee_id", "Contact.EmployeeID", "required", param)
		sl.ReportError(contact.CompanyID, "company_id", "Contact.CompanyID", "required", param)
	}
	param = "email or contact_no"
	if contact.Email == "" && contact.ContactNo == "" {
		sl.ReportError(contact.Email, "email", "Contact.Email", "required", param)
		sl.ReportError(contact.Email, "contact_no", "Contact.ContactNo", "required", param)
	}
	if contact.Email != "" {
		if contact.EmailTypeID == nil {
			sl.ReportError(contact.EmailTypeID, "email_type_id", "Contact.EmailTypeID", "required", "")
		}
	}
	if contact.ContactNo != "" {
		if contact.ContactTypeID == nil {
			sl.ReportError(contact.ContactTypeID, "contact_type_id", "Contact.ContactTypeID", "required", "")
		}
	}
}
