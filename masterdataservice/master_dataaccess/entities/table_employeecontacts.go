package entities

import "github.com/jinzhu/gorm"

type TableEmployeeContact struct {
	gorm.Model
	ContactId  uint `gorm:"column:contactid;not_null"`
	EmployeeId uint `gorm:"column:employeeid;not_null"`
	Contact    TableContacts
	Employee   TableEmployee
}

func (c TableEmployeeContact) TableName() string {
	return "table_employeecontact"
}
