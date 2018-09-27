package model

import (
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
)

type ServiceRpt struct {
	gorm.Model
	EmployeeID uint
	IssueDt    time.Time
	FromNo     uint
	ToNo       uint
}

func (u *ServiceRpt) Bind(r *http.Request) error {
	return nil
}
