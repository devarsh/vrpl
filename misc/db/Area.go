package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getAreaBy(condition ...interface{}) ([]*model.Area, error) {
	out := []*model.Area{}
	err := m.db.Select("id, name, city_id").Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) GetAreasByCityId(ctx context.Context, id uint) ([]*model.Area, error) {
	return m.getAreaBy("city_id = ?", id)
}

func (m *MiscDb) GetAreasByName(ctx context.Context, name string) ([]*model.Area, error) {
	return m.getAreaBy("name like '%?%'", name)
}

func (m *MiscDb) GetAllAreas(ctx context.Context) ([]*model.Area, error) {
	return m.getAreaBy()
}

func (m *MiscDb) GetAreaById(ctx context.Context, id uint) (*model.Area, error) {
	res, err := m.getAreaBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) AddArea(ctx context.Context, area *model.Area) (uint, error) {
	err := m.db.Create(area).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return area.ID, nil
}

func (m *MiscDb) UpdateArea(ctx context.Context, area *model.Area) error {
	column := map[string]interface{}{"name": area.Name, "city_id": area.CityID}
	err := m.db.Model(area).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteArea(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Area{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
