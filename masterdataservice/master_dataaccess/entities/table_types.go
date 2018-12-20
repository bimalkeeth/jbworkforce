package entities

import "github.com/jinzhu/gorm"

type TableTypes struct {
	gorm.Model
	TypeName          string            `gorm:"column:typename;not_null"`
	TypeAbbr          string            `gorm:"column:typeabbr;not_null"`
	TypeCategoryId    uint              `gorm:"column:typecategoryid;not_null"`
	TableTypeCategory TableTypeCategory `gorm:"foreignkey:typecategoryid"`
}

func (c TableTypes) TableName() string {
	return "table_types"
}
