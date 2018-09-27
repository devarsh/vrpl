package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
)

type Item struct {
	gorm.Model
	MachineTypeID uint   `gorm:"type:int(10) unsigned" validate:"required" json:"machine_type_id"`
	MachineType   Types  `validate:"-" json:"machine_type"`
	BrandName     string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"brand_name"`
	ModelName     string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"model_name"`
	HsnCode       string `gorm:"type:varchar(20)" validate:"required,lte=50" json:"hsn_code"`
}

func (u *Item) Bind(r *http.Request) error {
	return nil
}
