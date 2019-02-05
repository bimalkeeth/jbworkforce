package entities

import "github.com/jinzhu/gorm"

type TableEmployeeGroupEmp struct {
	gorm.Model
	EmployeeGroupId uint `gorm:"column:employeegroupid;not_null"`
	EmployeeId      uint `gorm:"column:employeeid;not_null"`
	EmployeeGroup   *TableEmployeeGroup
	Employee        *TableEmployee
}

func (c TableEmployeeGroupEmp) TableName() string {
	return "table_employeegroupemp"
}
