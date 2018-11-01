package entities

import "github.com/jinzhu/gorm"

type TableDepartment struct {
	gorm.Model
	DepartmentName string `gorm:"column:departmentname;not_null"`
	DepartmentAbbr string `gorm:"column:departmentabbr;not_null"`
	Description    string `gorm:"column:description"`
	Definition     string `gorm:"column:definition"`
	IsActive       bool   `gorm:"column:isactive;not_null;default:true"`
}

func (c TableDepartment) TableName() string {
	return "table_department"
}
