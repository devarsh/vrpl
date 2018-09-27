package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getItemBy(condition ...interface{}) ([]*model.Item, error) {
	out := []*model.Item{}
	err := m.db.
		Preload("MachineType").
		Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) AddItem(ctx context.Context, item *model.Item) (uint, error) {
	err := m.db.Create(item).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return item.ID, nil
}

func (m *MiscDb) UpdateItem(ctx context.Context, item *model.Item) error {
	column := map[string]interface{}{
		"machine_type_id": item.MachineTypeID,
		"brand_name":      item.BrandName,
		"model_name":      item.ModelName,
		"hsn_code":        item.HsnCode,
	}
	err := m.db.Model(item).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteItem(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Item{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) GetItemByID(ctx context.Context, id uint) (*model.Item, error) {
	res, err := m.getItemBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) GetAllItems(ctx context.Context) ([]*model.Item, error) {
	return m.getItemBy()
}
