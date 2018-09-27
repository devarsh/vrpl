package jwt

import (
	"context"
	customError "github.com/devarsh/vrpl/error"
	"github.com/devarsh/vrpl/jwt/model"
	"github.com/jinzhu/gorm"
)

type JwtDB struct {
	db *gorm.DB
}

var mySQLErr = customError.ErrMySqlDBError

func (m *JwtDB) getBy(condition ...interface{}) (*model.JwtPersist, error) {
	out := model.JwtPersist{}
	err := m.db.Select("id, unique_id, token, expires_at, user_id, blacklist").First(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return &out, nil
}

func (m *JwtDB) deleteBy(query interface{}, args ...interface{}) (bool, error) {
	err := m.db.Where(query, args).Delete(&model.JwtPersist{}).Error
	if err != nil {
		return false, mySQLErr(err)
	}
	return true, nil
}

func (m *JwtDB) Add(ctx context.Context, token *model.JwtPersist) (uint, error) {
	err := m.db.Create(token).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return token.ID, nil
}

func (m *JwtDB) GetByUUID(ctx context.Context, uuid string) (*model.JwtPersist, error) {
	return m.getBy("unique_id = ?", uuid)
}

func (m *JwtDB) GetByUserID(ctx context.Context, userID uint) ([]*model.JwtPersist, error) {
	tokens := make([]*model.JwtPersist, 0)
	err := m.db.Select("id, unique_id, token, expires_at, user_id, blacklist").Where("user_id = ?", userID).Find(&tokens).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return tokens, nil
}

func (m *JwtDB) DeleteByUUID(ctx context.Context, uuid string) (bool, error) {
	return m.deleteBy("unique_id = ?", uuid)
}

func (m *JwtDB) DeleteByUserID(ctx context.Context, userID int) (bool, error) {
	return m.deleteBy("user_id = ?", userID)
}

func (m *JwtDB) Update(ctx context.Context, token *model.JwtPersist) (bool, error) {
	columns := map[string]interface{}{"blacklist": token.Blacklist}
	err := m.db.Model(token).Update(columns).Error
	if err != nil {
		return false, mySQLErr(err)
	}
	return true, nil
}

func (m *JwtDB) UpdateByUserID(ctx context.Context, userID uint) (bool, error) {
	columns := map[string]interface{}{"blacklist": 0}
	err := m.db.Model(&model.JwtPersist{}).Update(columns).Where("user_id", userID).Error
	if err != nil {
		return false, mySQLErr(err)
	}
	return true, nil
}
