package entities

import "github.com/jinzhu/gorm"

type TableSmslogType struct {
	gorm.Model
	TypeName string `gorm:"column:typename;not_null"`
	TypeAbbr string `gorm:"column:typeabbr;not_null"`
}

func (c TableSmslogType) TableName() string {
	return "table_smslogtype"
}
