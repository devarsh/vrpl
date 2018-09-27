package db

import (
	customError "github.com/devarsh/vrpl/error"
	"github.com/jinzhu/gorm"
)

type MasterDb struct {
	db *gorm.DB
}

var (
	mySQLErr      = customError.ErrMySqlDBError
	mySQLNotFound = customError.DBRecordNotFound()
)

func NewMasterDb(mydb *gorm.DB) *MasterDb {
	return &MasterDb{db: mydb}
}
