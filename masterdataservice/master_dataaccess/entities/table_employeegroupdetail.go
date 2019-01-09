package entities

import "github.com/jinzhu/gorm"

type TableEmployeeGroupDetail struct {
	gorm.Model
	EmployeeGroupId uint `gorm:"column:employeegroupid;not_null"`
	EmployeeId      uint `gorm:"column:employeeid;not_null"`
	EmployeeGroup   TableEmployeeGroup
	Employee        TableEmployee
}

func (c TableEmployeeGroupDetail) TableName() string {
	return "table_employeegroupdetail"
}
