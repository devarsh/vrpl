package db

import (
	"context"
	"github.com/devarsh/vrpl/master/model"
)

func (m *MasterDb) getMachineBy(condition ...interface{}) ([]*model.Machine, error) {
	out := []*model.Machine{}
	err := m.db.
		Preload("Item").
		Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MasterDb) GetAllMachines() ([]*model.Machine, error) {
	return m.getMachineBy()
}

func (m *MasterDb) GetMachinesByEmployeeID(id uint) ([]*model.Machine, error) {
	return m.getMachineBy("employee_id = ?", id)
}

func (m *MasterDb) GetMachinesByClientID(id uint) ([]*model.Machine, error) {
	return m.getMachineBy("client_id = ?", id)
}

func (m *MasterDb) GetMachineByID(id uint) (*model.Machine, error) {
	res, err := m.getMachineBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MasterDb) AddMachine(ctx context.Context, machine *model.Machine) (uint, error) {
	err := m.db.Create(machine).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return machine.ID, nil
}

func (m *MasterDb) UpdateMachine(ctx context.Context, machine *model.Machine) error {
	column := map[string]interface{}{
		"serial_no": machine.SerialNo, "remarks": machine.Remarks, "to_dt": machine.EmployeeID,
	}
	err := m.db.Model(machine).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MasterDb) DeleteMachine(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Machine{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
