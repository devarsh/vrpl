package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getCityBy(condition ...interface{}) ([]*model.City, error) {
	out := []*model.City{}
	err := m.db.Select("id, name, state_id").Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) GetCitiesByStateId(ctx context.Context, id uint) ([]*model.City, error) {
	return m.getCityBy("state_id = ?", id)
}

func (m *MiscDb) GetCitiesByName(ctx context.Context, name string) ([]*model.City, error) {
	return m.getCityBy("name like '%?%'", name)
}

func (m *MiscDb) GetAllCities(ctx context.Context) ([]*model.City, error) {
	return m.getCityBy()
}

func (m *MiscDb) GetCityByID(ctx context.Context, id uint) (*model.City, error) {
	res, err := m.getCityBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) AddCity(ctx context.Context, city *model.City) (uint, error) {
	err := m.db.Create(city).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return city.ID, nil
}

func (m *MiscDb) UpdateCity(ctx context.Context, city *model.City) error {
	column := map[string]interface{}{"name": city.Name, "city_id": city.StateID}
	err := m.db.Model(city).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteCity(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.City{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
