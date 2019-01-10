package entities

import "github.com/jinzhu/gorm"

type TableSubDepartment struct {
	gorm.Model
	DepartmentId      uint   `gorm:"column:departmentid;not_null"`
	CostCentreId      uint   `gorm:"column:costcentreid;not_null"`
	SubDepartmentName string `gorm:"column:subdepartmentname;type:varchar(200);not_null"`
	SubDepartmentAttr string `gorm:"column:subdepartmentattr;type:varchar(50);not_null"`
	Description       string `gorm:"column:description;type:varchar(400);not_null"`
	IsActive          bool   `gorm:"column:isactive;not_null"`
	Department        TableDepartment
	CostCentre        TableCostCentre
}

func (c TableSubDepartment) TableName() string {
	return "table_subdepartment"
}
