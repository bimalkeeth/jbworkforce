package entities

import "github.com/jinzhu/gorm"

type TableShiftType struct {
	gorm.Model
	ShiftTypeName string `gorm:"column:shifttypename;type:varchar(200);not_null"`
	ShiftTypeAbbr string `gorm:"column:shifttypeabbr;type:varchar(50);not_null"`
}

func (c TableShiftType) TableName() string {
	return "table_shifttype"
}
