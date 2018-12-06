package entities

import "github.com/jinzhu/gorm"

type TableEmployeeContact struct {
	gorm.Model
	ContactId     uint          `gorm:"column:contactid;not_null"`
	EmployeeId    uint          `gorm:"column:employeeid;not_null"`
	TableContact  TableContacts `gorm:"foreignkey:contactid"`
	TableEmployee TableEmployee `gorm:"foreignkey:employeeid"`
}

func (c TableEmployeeContact) TableName() string {
	return "table_employeecontact"
}
