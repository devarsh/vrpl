package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getCountryBy(condition ...interface{}) ([]*model.Country, error) {
	out := []*model.Country{}
	err := m.db.Select("id, name").Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) GetCountriesByName(ctx context.Context, name string) ([]*model.Country, error) {
	return m.getCountryBy("name like '%?%'", name)
}

func (m *MiscDb) GetAllCountries(ctx context.Context) ([]*model.Country, error) {
	return m.getCountryBy()
}

func (m *MiscDb) GetCountryByID(ctx context.Context, id uint) (*model.Country, error) {
	res, err := m.getCountryBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) AddCountry(ctx context.Context, country *model.Country) (uint, error) {
	err := m.db.Create(country).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return country.ID, nil
}

func (m *MiscDb) UpdateCountry(ctx context.Context, country *model.Country) error {
	column := map[string]interface{}{"name": country.Name}
	err := m.db.Model(country).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteCountry(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Country{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
