package db

import (
	"context"
	"github.com/devarsh/vrpl/misc/model"
)

func (m *MiscDb) getContactBy(condition ...interface{}) ([]*model.Contact, error) {
	out := []*model.Contact{}
	err := m.db.
		Preload("ContactType").
		Preload("EmailType").
		Find(&out, condition...).Error
	if err != nil {
		return nil, mySQLErr(err)
	}
	return out, nil
}

func (m *MiscDb) AddContact(ctx context.Context, contact *model.Contact) (uint, error) {
	err := m.db.Create(contact).Error
	if err != nil {
		return 0, mySQLErr(err)
	}
	return contact.ID, nil
}

func (m *MiscDb) UpdateContact(ctx context.Context, contact *model.Contact) error {
	column := map[string]interface{}{
		"contact_no":      contact.ContactNo,
		"contact_person":  contact.ContactPerson,
		"contact_type_id": contact.ContactTypeID,
		"designation":     contact.Designation,
		"email":           contact.Email,
		"email_type_id":   contact.EmailTypeID,
	}
	err := m.db.Model(contact).Update(column).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) DeleteContact(ctx context.Context, id uint) error {
	err := m.db.Where("id = ?", id).Delete(&model.Contact{}).Error
	if err != nil {
		return mySQLErr(err)
	}
	return nil
}

func (m *MiscDb) GetContactById(ctx context.Context, id uint) (*model.Contact, error) {
	res, err := m.getContactBy("id = ?", id)
	if err != nil {
		return nil, mySQLErr(err)
	}
	if len(res) <= 0 {
		return nil, mySQLNotFound
	}
	return res[0], nil
}

func (m *MiscDb) GetContactByClientID(ctx context.Context, id uint) ([]*model.Contact, error) {
	return m.getContactBy("client_id = ?", id)

}

func (m *MiscDb) GetContactByEmployeeID(ctx context.Context, id uint) ([]*model.Contact, error) {
	return m.getContactBy("employee_id = ?", id)
}

func (m *MiscDb) GetContactByCompanyID(ctx context.Context, id uint) ([]*model.Contact, error) {
	return m.getContactBy("company_id = ?", id)
}
