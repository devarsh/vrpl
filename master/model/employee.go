package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Employee struct {
	gorm.Model
	Name           string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
	ContractStatus uint8  `gorm:"type:tinyint(1)" json:"contract_status"`
	PanNo          string `gorm:"type:varchar(20)" validate:"lte=20" json:"pan_no"`
	BankName       string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"bank_name"`
	BranchName     string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"branch_name"`
	IfsCode        string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"ifs_code"`
	AccountNo      string `gorm:"type:varchar(20)" validate:"required,lte=20" json:"account_no"`
}

func (u *Employee) Bind(r *http.Request) error {
	return nil
}
