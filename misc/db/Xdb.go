package db

import (
	customError "github.com/devarsh/vrpl/error"
	"github.com/jinzhu/gorm"
)

type MiscDb struct {
	db *gorm.DB
}

var (
	mySQLErr      = customError.ErrMySqlDBError
	mySQLNotFound = customError.DBRecordNotFound()
)

func NewMiscDb(mydb *gorm.DB) *MiscDb {
	return &MiscDb{db: mydb}
}
