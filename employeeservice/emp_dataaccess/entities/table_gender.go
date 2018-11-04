package entities

import "github.com/jinzhu/gorm"

type TableGender struct {
	gorm.Model
	Name string `gorm:"column:name;not_null"`
}

func (c TableGender) TableName() string {
	return "table_gender"
}
