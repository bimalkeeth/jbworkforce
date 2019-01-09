package entities

import "github.com/jinzhu/gorm"

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
