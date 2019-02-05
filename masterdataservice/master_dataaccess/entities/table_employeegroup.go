package entities

import "github.com/jinzhu/gorm"

type TableEmployeeGroup struct {
	gorm.Model
	EmployeeGroupName string `gorm:"column:employeegroupname;not_null"`
	Description       string `gorm:"column:description"`
	EmployeeGroupEmps []*TableEmployeeGroupEmp
}

func (c TableEmployeeGroup) TableName() string {
	return "table_employeegroup"
}
