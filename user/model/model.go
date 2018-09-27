package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"firstName"`
	LastName  string `gorm:"type:varchar(50)" validate:"required,lte=50" json:"lastName"`
	Email     string `gorm:"type:varchar(50);unique_index" validate:"required,let=50,email" json:"email"`
	Password  string `gorm:"type:varchar(70)" validate:"required,gte=8,lte=50" json:"password"`
	Status    uint8  `gorm:"type:tinyint(1)" json:"status"`
}

func (u *User) HashPassword() bool {
	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return false
	}
	u.Password = string(h)
	return true
}

func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

type UserLogin struct {
	Username string `validate:"required,email" json:"username"`
	Password string `validate:"required,gte=8,lte=50" json:"password"`
}

type UserPasswordChange struct {
	Username    string `validate:"required,email" json:"username"`
	Password    string `validate:"omitempty,gte=8,lte=50" json:"password"`
	NewPassword string `validate:"required,gte=8,lte=50" json:"newPassword"`
}

type UserStatus struct {
	Username string `validate:"required,email" json:"username"`
}
