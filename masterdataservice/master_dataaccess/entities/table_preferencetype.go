package entities

import "github.com/jinzhu/gorm"

type TablePreferenceType struct {
	gorm.Model
	PreferenceName string `gorm:"column:preferencename"`
}

func (c TablePreferenceType) TableName() string {
	return "table_preferencetype"
}
