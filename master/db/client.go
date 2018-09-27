package db

import (
	"context"
	"fmt"
	"github.com/devarsh/vrpl/master/model"
)

func (m *MasterDb) getClientBy(condition ...interface{}) ([]*model.Client, error) {
	out := []*model.Client{}
	err := m.db.
		Preload("Group").Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MasterDb) GetClientsByName(name string) ([]*model.Client, error) {
	fmt.Println(name)
	return m.getClientBy("client.name like ?", fmt.Sprintf("%%%s%%", name))
}

func (m *MasterDb) GetClientsByTallyName(name string) ([]*model.Client, error) {
	return m.getClientBy("client.tally_name like ?", fmt.Sprintf("%%%s%%", name))
}

func (m *MasterDb) GetClientByID(id uint) (*model.Client, error) {
	res, err := m.getClientBy("client.id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MasterDb) GetClientByGroupID(id uint) ([]*model.Client, error) {
	return m.getClientBy("client.group_id = ?", id)
}

func (m *MasterDb) AddClient(ctx context.Context, client *model.Client) (uint, error) {
	err := m.db.Create(client).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return client.ID, nil
}

func (m *MasterDb) UpdateClient(ctx context.Context, client *model.Client) error {
	column := map[string]interface{}{
		"branch_code": client.BranchCode, "name": client.Name,
		"tags": client.Tags, "group_id": client.GroupID,
	}
	err := m.db.Model(client).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MasterDb) DeleteClient(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Client{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}
