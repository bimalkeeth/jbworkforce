package entities

import "github.com/jinzhu/gorm"

type TableContactType struct {
	gorm.Model
	ContactType string `gorm:"column:contacttype;not_null"`
	Description string `gorm:"column:description"`
}

func (c TableContactType) TableName() string {
	return "table_contacttype"
}
