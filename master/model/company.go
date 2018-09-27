package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Company struct {
	gorm.Model
	Name       string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
	ShortName  string `gorm:"type:varchar(60)" validate:"lte=60" json:"short_name"`
	PanNo      string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"pan_no"`
	Gstin      string `gorm:"type:varchar(20)" validate:"lte=20" json:"gstin"`
	BankName   string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"bank_name"`
	BranchName string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"branch_name"`
	IfsCode    string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"ifs_code"`
	AccountNo  string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"account_no"`
}

func (u *Company) Bind(r *http.Request) error {
	return nil
}
