package entities

import "github.com/jinzhu/gorm"

type TableTypeCategory struct {
	gorm.Model
	TypeCategoryName string `gorm:"column:typecategoryname;not_null"`
	TypeCategoryAbbr string `gorm:"column:typecategoryabbr;not_null"`
}

func (c TableTypeCategory) TableName() string {
	return "table_typecategory"
}
