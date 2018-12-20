package entities

import "github.com/jinzhu/gorm"

type TableStatusType struct {
	gorm.Model
	StatusTypeName string        `gorm:"column:statustypename;not_null"`
	StatusTypeAbbr string        `gorm:"column:statustypeabbr;not_null"`
	TableStatus    []TableStatus `gorm:"foreignkey:statustypeid"`
}

func (c TableStatusType) TableName() string {
	return "table_statustype"
}
