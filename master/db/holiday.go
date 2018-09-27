package db

import (
	"context"
	"github.com/devarsh/vrpl/master/model"
	"time"
)

func (m *MasterDb) getHolidayBy(condition ...interface{}) ([]*model.Holiday, error) {
	out := []*model.Holiday{}
	err := m.db.Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MasterDb) GetAllHolidays() ([]*model.Holiday, error) {
	return m.getHolidayBy()
}

func (m *MasterDb) GetHolidayByID(id uint) (*model.Holiday, error) {
	res, err := m.getHolidayBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MasterDb) GetHolidayByDate(date time.Time) ([]*model.Holiday, error) {
	return m.getHolidayBy("date = ? ", date)
}

func (m *MasterDb) AddHoliday(ctx context.Context, holiday *model.Holiday) (uint, error) {
	err := m.db.Create(holiday).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return holiday.ID, nil
}

func (m *MasterDb) UpdateHoliday(ctx context.Context, holiday *model.Holiday) error {
	column := map[string]interface{}{
		"name": holiday.Name, "date": holiday.Date,
	}
	err := m.db.Model(holiday).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MasterDb) DeleteHoliday(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Holiday{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
