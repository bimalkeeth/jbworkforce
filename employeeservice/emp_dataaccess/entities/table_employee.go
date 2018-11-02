package entities

import "github.com/jinzhu/gorm"

type TableEmployee struct {
	gorm.Model
	addressid uint
	firstname string
	lastname  string
}

func (c TableEmployee) TableName() string {
	return "table_employee"
}
