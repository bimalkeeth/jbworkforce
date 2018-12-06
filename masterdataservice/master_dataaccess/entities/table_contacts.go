package entities

import "github.com/jinzhu/gorm"

type TableContacts struct {
	gorm.Model
	ContactTypeId    uint                   `gorm:"column:contacttypeid;not_null"`
	Contact          string                 `gorm:"column:contact;not_null"`
	AgencyContacts   []TableAgencyContacts  `gorm:"foreignkey:contactid"`
	EmployeeContact  []TableEmployeeContact `gorm:"foreignkey:contactid"`
	TableContactType TableContactType       `gorm:"foreignkey:contacttypeid"`
}

func (c TableContacts) TableName() string {
	return "table_contacts"
}
