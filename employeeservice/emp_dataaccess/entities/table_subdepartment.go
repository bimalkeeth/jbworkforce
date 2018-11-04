package entities

import "github.com/jinzhu/gorm"

type TableSubDepartment struct {
	gorm.Model
}

func (c TableSubDepartment) TableName() string {
	return "table_subdepartment"
}
