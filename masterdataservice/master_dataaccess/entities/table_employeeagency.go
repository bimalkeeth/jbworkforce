package entities

import "github.com/jinzhu/gorm"

type TableEmployeeAgency struct {
	gorm.Model
	EmployeeId uint `gorm:"column:employeeid;not_null"`
	AgencyId   uint `gorm:"column:agencyid;not_null"`
	Employee   TableEmployee
	Agency     TableAgency
}

func (c TableEmployeeAgency) TableName() string {
	return "table_employeeagency"
}
