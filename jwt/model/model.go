package model

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"time"
)

type JwtPersist struct {
	gorm.Model
	UniqueID  string    `gorm:"type:char(36);unique_index"`
	Token     string    `gorm:"type:varchar(2000)" validate:"required"`
	ExpiresAt time.Time `gorm:"type:timestamp"`
	UserId    uint      `gorm:"type:int(10) unsigned"`
	Blacklist uint8     `gorm:"type:tinyint(1)"`
}

type TokenPayload struct {
	UserId uint   `json:"user_id"`
	Roles  string `json:"roles"`
	*jwt.StandardClaims
}
