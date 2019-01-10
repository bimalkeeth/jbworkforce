package entities

import "github.com/jinzhu/gorm"

type TableStatusType struct {
	gorm.Model
	StatusTypeName string `gorm:"column:statustypename;type:varchar(200);not_null"`
	StatusTypeAbbr string `gorm:"column:statustypeabbr;type:varchar(50);not_null"`
	Status         []TableStatus
}

func (c TableStatusType) TableName() string {
	return "table_statustype"
}
