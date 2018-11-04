package entities

import "github.com/jinzhu/gorm"

type TableEmployeeType struct {
	gorm.Model
	EmployeeTypeName string `gorm:"column:employeetypename;not_null"`
	EmployeeTypeAttr string `gorm:"column:employeetypeattr;not_null"`
	DisplayOrder     int    `gorm:"column:qualificationtypename;not_null;Default:0"`
}

func (c TableEmployeeType) TableName() string {
	return "table_employeetype"
}
