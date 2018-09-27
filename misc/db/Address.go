package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getAddressBy(condition ...interface{}) ([]*model.Address, error) {
	out := []*model.Address{}
	err := m.db.
		Preload("Country").
		Preload("State").
		Preload("City").
		Preload("Area").
		Preload("AddressType").
		Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) AddAddress(ctx context.Context, address *model.Address) (uint, error) {
	err := m.db.Create(address).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return address.ID, nil
}

func (m *MiscDb) UpdateAddress(ctx context.Context, address *model.Address) error {
	column := map[string]interface{}{
		"line1": address.Line1, "line2": address.Line2, "line3": address.Line3,
		"line4": address.Line4, "pincode": address.Pincode, "area_id": address.AreaID,
		"city_id": address.CityID, "state_id": address.StateID, "country_id": address.CountryID,
		"address_type_id": address.AddressTypeID,
	}
	err := m.db.Model(address).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteAddress(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Area{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) GetAddressById(ctx context.Context, id uint) (*model.Address, error) {
	res, err := m.getAddressBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) GetAddressByClientID(ctx context.Context, id uint) ([]*model.Address, error) {
	return m.getAddressBy("client_id = ?", id)

}

func (m *MiscDb) GetAddressByCompanyID(ctx context.Context, id uint) ([]*model.Address, error) {
	return m.getAddressBy("company_id = ?", id)
}

func (m *MiscDb) GetAddressByEmployeeID(ctx context.Context, id uint) ([]*model.Address, error) {
	return m.getAddressBy("employee_id = ?", id)
}
