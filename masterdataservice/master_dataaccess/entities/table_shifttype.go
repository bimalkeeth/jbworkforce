package entities

import "github.com/jinzhu/gorm"

type TableShiftType struct {
	gorm.Model
	ShiftTypeName string `gorm:"column:shifttypename;not_null"`
	ShiftTypeAbbr string `gorm:"column:shifttypeabbr;not_null"`
}

func (c TableShiftType) TableName() string {
	return "table_shifttype"
}
