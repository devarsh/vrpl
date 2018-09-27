package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getStateBy(condition ...interface{}) ([]*model.State, error) {
	out := []*model.State{}
	err := m.db.Select("id, name, country_id").Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) GetStatesByCountryId(ctx context.Context, id uint) ([]*model.State, error) {
	return m.getStateBy("country_id = ?", id)
}

func (m *MiscDb) GetStatesByName(ctx context.Context, name string) ([]*model.State, error) {
	return m.getStateBy("name like '%?%'", name)
}

func (m *MiscDb) GetAllStates(ctx context.Context) ([]*model.State, error) {
	return m.getStateBy()
}

func (m *MiscDb) GetStateById(ctx context.Context, id uint) (*model.State, error) {
	res, err := m.getStateBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) AddState(ctx context.Context, state *model.State) (uint, error) {
	err := m.db.Create(state).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return state.ID, nil
}

func (m *MiscDb) UpdateState(ctx context.Context, state *model.State) error {
	column := map[string]interface{}{"name": state.Name, "city_id": state.CountryID}
	err := m.db.Model(state).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteState(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.State{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
