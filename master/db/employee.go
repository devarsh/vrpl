package db

import (
	"context"
	"github.com/devarsh/vrpl/master/model"
)

func (m *MasterDb) getEmployeeBy(condition ...interface{}) ([]*model.Employee, error) {
	out := []*model.Employee{}
	err := m.db.Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MasterDb) GetAllEmployees() ([]*model.Employee, error) {
	return m.getEmployeeBy()
}

func (m *MasterDb) GetEmployeeByID(id uint) (*model.Employee, error) {
	res, err := m.getEmployeeBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MasterDb) AddEmployee(ctx context.Context, employee *model.Employee) (uint, error) {
	err := m.db.Create(employee).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return employee.ID, nil
}

func (m *MasterDb) UpdateEmployee(ctx context.Context, employee *model.Employee) error {
	column := map[string]interface{}{
		"name": employee.Name, "contract_status": employee.ContractStatus,
		"pan_no": employee.PanNo, "bank_name": employee.BankName,
		"branch_name": employee.BranchName, "ifs_code": employee.IfsCode,
		"account_no": employee.AccountNo,
	}
	err := m.db.Model(employee).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MasterDb) DeleteEmployee(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Employee{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
