package entities

import "github.com/jinzhu/gorm"

type TableSmslogType struct {
	gorm.Model
	TypeName string `gorm:"column:typename;type:varchar(200);not_null"`
	TypeAbbr string `gorm:"column:typeabbr;type:varchar(50);not_null"`
	Smslogs  []TableSmsLog
}

func (c TableSmslogType) TableName() string {
	return "table_smslogtype"
}
