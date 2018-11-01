package entities

import "github.com/jinzhu/gorm"

type TableEmployee struct {
	gorm.Model
}

func (c TableEmployee) TableName() string {
	return "table_employee"
}
