package entities

import "github.com/jinzhu/gorm"

type TableAllowanceLevel struct {
	gorm.Model
	AllowanceLevelName string `gorm:"column:allowancelevelname;type:varchar(50);not_null"`
	Description        string `gorm:"column:description;type:varchar(400)"`
}

func (c TableAllowanceLevel) TableName() string {
	return "table_allowancelevel"
}
