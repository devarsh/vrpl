package db

import (
	"context"
	"github.com/devarsh/vrpl/master/model"
)

func (m *MasterDb) getCompanyBy(condition ...interface{}) ([]*model.Company, error) {
	out := []*model.Company{}
	err := m.db.Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MasterDb) GetAllCompanies() ([]*model.Company, error) {
	return m.getCompanyBy()
}

func (m *MasterDb) GetCompanyByID(id uint) (*model.Company, error) {
	res, err := m.getCompanyBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MasterDb) AddCompany(ctx context.Context, company *model.Company) (uint, error) {
	err := m.db.Create(company).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return company.ID, nil
}

func (m *MasterDb) UpdateCompany(ctx context.Context, company *model.Company) error {
	column := map[string]interface{}{
		"name": company.Name, "short_name": company.ShortName,
		"pan_no": company.PanNo, "gstin": company.Gstin,
	}
	err := m.db.Model(company).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MasterDb) DeleteCompany(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Company{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
