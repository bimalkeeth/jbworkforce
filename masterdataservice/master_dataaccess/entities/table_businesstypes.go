package entities

import "github.com/jinzhu/gorm"

type TableBusinessTypes struct {
	gorm.Model
	BusinessTypeName string          `gorm:"column:statusname;not_null"`
	BusinessTypeAbbr string          `gorm:"column:statusabbr;not_null"`
	TypeCategoryId   uint            `gorm:"column:statustypeid;not_null"`
	TypeCategory     TableStatusType `gorm:"foreignkey:statustypeid"`
}

func (c TableBusinessTypes) TableName() string {
	return "table_businesstypes"
}
