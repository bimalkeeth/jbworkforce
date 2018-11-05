package entities

import "github.com/jinzhu/gorm"

type TableAllowanceLevel struct {
	gorm.Model
	AllowanceLevelName string `gorm:"column:allowancelevelname;not_null"`
}

func (c TableAllowanceLevel) TableName() string {
	return "table_allowancelevel"
}
