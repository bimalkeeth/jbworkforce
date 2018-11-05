package entities

import "github.com/jinzhu/gorm"

type TableSubDepartment struct {
	gorm.Model
	DepartmentId      uint   `gorm:"column:departmentid;not_null"`
	CostCentreId      uint   `gorm:"column:costcentreid;not_null"`
	SubDepartmentName string `gorm:"column:subdepartmentname;not_null"`
	SubDepartmentAttr string `gorm:"column:subdepartmentattr;not_null"`
	Description       string `gorm:"column:description;not_null"`
	IsActive          bool   `gorm:"column:isactive;not_null"`

	TableDepartment TableDepartment `gorm:"foreignkey:departmentid"`
	TableCostCentre TableCostCentre `gorm:"foreignkey:costcentreid"`
}

func (c TableSubDepartment) TableName() string {
	return "table_subdepartment"
}
