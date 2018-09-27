package user

import (
	"context"
	"fmt"
	customErr "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/user/model"
	"github.com/jinzhu/gorm"
)

type UserManager struct {
	db *UserDB
}

func NewUserManager(db *gorm.DB) *UserManager {
	if db == nil {
		panic("Nil Db instance passed")
	}
	userdb := &UserDB{db: db}
	return &UserManager{db: userdb}
}

func (us *UserManager) AddUser(data *model.User) (uint, error) {
	ctx := context.Background()
	ok := data.HashPassword()
	data.Status = 0
	if !ok {
		return 0, customErr.ErrInvalidRequestError(fmt.Errorf("Error Generating Hash"))
	}
	id, err := us.db.Add(ctx, data)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (us *UserManager) VerifyPassword(data *model.UserLogin) (*model.User, error) {
	ctx := context.Background()
	res, err := us.db.GetByEmail(ctx, data.Username, true)
	if err != nil {
		return nil, err
	}
	if res.Status == 1 {
		return nil, customErr.UserDisable()
	}
	ok := res.CheckPassword(data.Password)
	if !ok {
		return nil, customErr.UserPassInvalid()
	}
	res.Password = ""
	return res, nil
}

func (us *UserManager) ChangePassword(data *model.UserPasswordChange, force bool) error {
	ctx := context.Background()
	res, err := us.db.GetByEmail(ctx, data.Username, true)
	if err != nil {
		return err
	}
	if !force {
		ok := res.CheckPassword(data.Password)
		if !ok {
			return customErr.UserPassInvalid()
		}
	}
	res.Password = data.NewPassword
	res.HashPassword()
	err = us.db.Update(ctx, res, true)
	if err != nil {
		return err
	}
	return nil
}
