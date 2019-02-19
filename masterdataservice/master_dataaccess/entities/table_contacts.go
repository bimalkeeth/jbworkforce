package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableContacts struct {
	gorm.Model
	ContactTypeId   uint   `gorm:"column:contacttypeid;not_null"`
	Contact         string `gorm:"column:contact;type:varchar(250);not_null"`
	AgencyContacts  []TableAgencyContacts
	EmployeeContact []TableEmployeeContact
	ContactType     TableContactType
}

func (c TableContacts) TableName() string {
	return "table_contacts"
}

func (c TableContacts) Validate(db *gorm.DB) {
	if c.ContactTypeId == 0 {
		_ = db.AddError(errors.New("contact type should should not be empty"))
	}
	if len(c.Contact) > 250 {
		_ = db.AddError(errors.New("contact contains more than 250 characters"))
	}
}
