package entities

import "github.com/jinzhu/gorm"

type TableEmployeeType struct {
	gorm.Model
	EmployeeTypeName string `gorm:"column:employeetypename;type:varchar(150);not_null"`
	EmployeeTypeAttr string `gorm:"column:employeetypeattr;type:varchar(150);not_null"`
	DisplayOrder     int    `gorm:"column:qualificationtypename;not_null;Default:0"`
}

func (c TableEmployeeType) TableName() string {
	return "table_employeetype"
}
