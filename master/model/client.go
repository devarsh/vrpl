package model

import (
	"github.com/devarsh/vrpl/misc/model"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Client struct {
	gorm.Model
	GroupID    uint        `gorm:"type:int(10) unsigned" validate:"required" json:"group_id"`
	Group      model.Types `validate:"-" json:"group"`
	Name       string      `gorm:"type:varchar(50)" validate:"required,lte=50" json:"name"`
	TallyName  string      `gorm:"type:varchar(60)" validate:"lte=60" json:"tally_name"`
	BranchCode string      `gorm:"type:varchar(50)" validate:"required,lte=50" json:"branch_code"`
	Tags       string      `gorm:"type:varchar(100)" validate:"lte=100" json:"tags"`
	PanNo      string      `gorm:"type:varchar(20)" validate:"lte=20" json:"pan_no"`
	Gstin      string      `gorm:"type:varchar(20)" validate:"lte=20" json:"gstin"`
}

func (u *Client) Bind(r *http.Request) error {
	return nil
}
