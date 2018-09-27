package user

import (
	"context"

	customError "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/user/model"
	"github.com/jinzhu/gorm"
)

type UserDB struct {
	db *gorm.DB
}

var mySQLErr = customError.ErrMySqlDBError

func (m *UserDB) getBy(fetchPwd bool, condition ...interface{}) (*model.User, error) {
	out := model.User{}
	var err error
	if fetchPwd {
		err = m.db.Select("id, first_name, last_name, email, password, status, created_at").First(&out, condition...).Error
	} else {
		err = m.db.Select("id, first_name, last_name, email, status, created_at").First(&out, condition...).Error
	}
	if err != nil {
		return nil, mySQLErr(err)
	}
	return &out, nil
}

func (m *UserDB) GetByID(ctx context.Context, id uint, fetchPwd bool) (*model.User, error) {
	return m.getBy(fetchPwd, "id = ?", id)
}

func (m *UserDB) GetByEmail(ctx context.Context, email string, fetchPwd bool) (*model.User, error) {
	return m.getBy(fetchPwd, "email = ?", email)
}

func (m *UserDB) Add(ctx context.Context, user *model.User) (uint, error) {
	err := m.db.Create(user).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return user.ID, nil
}

func (m *UserDB) Update(ctx context.Context, user *model.User, updatePwd bool) error {
	columns := map[string]interface{}{"first_name": user.FirstName, "last_name": user.LastName, "email": user.Email, "status": user.Status}
	if updatePwd {
		columns["password"] = user.Password
	}
	err := m.db.Model(user).Updates(columns).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *UserDB) Delete(ctx context.Context, id uint) error {
	err := m.db.Where("id = ? ", id).Delete(&model.User{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
