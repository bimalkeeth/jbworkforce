package entities

import "github.com/jinzhu/gorm"

type TableDepartment struct {
	gorm.Model
	DepartmentName string `gorm:"column:departmentname;type:varchar(200);not_null"`
	DepartmentAbbr string `gorm:"column:departmentabbr;type:varchar(50);not_null"`
	Description    string `gorm:"column:description;type:varchar(250);"`
	Definition     string `gorm:"column:definition;type:varchar(250);"`
	IsActive       bool   `gorm:"column:isactive;not_null;default:true"`
	DivisionId     uint   `gorm:"column:divisionid;not_null"`
	Division       TableDivision
	SubDepartments []TableSubDepartment
}

func (c TableDepartment) TableName() string {
	return "table_department"
}
