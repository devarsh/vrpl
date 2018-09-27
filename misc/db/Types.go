package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getTypeBy(condition ...interface{}) ([]*model.Types, error) {
	out := []*model.Types{}
	err := m.db.Select("id, group_nm, name").Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) GetTypesByGroupNm(ctx context.Context, group string) ([]*model.Types, error) {
	return m.getTypeBy("group_nm = ?", group)
}

func (m *MiscDb) GetAllTypes(ctx context.Context) ([]*model.Types, error) {
	return m.getTypeBy()
}

func (m *MiscDb) GetTypeById(ctx context.Context, id uint) (*model.Types, error) {
	res, err := m.getTypeBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) AddType(ctx context.Context, types *model.Types) (uint, error) {
	err := m.db.Create(types).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return types.ID, nil
}

func (m *MiscDb) UpdateType(ctx context.Context, types *model.Types) error {
	column := map[string]interface{}{"name": types.Name, "short_name": types.ShortName}
	err := m.db.Model(types).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteType(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Types{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
