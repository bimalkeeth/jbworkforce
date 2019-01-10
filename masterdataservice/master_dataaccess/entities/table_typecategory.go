package entities

import "github.com/jinzhu/gorm"

type TableTypeCategory struct {
	gorm.Model
	TypeCategoryName string `gorm:"column:typecategoryname;type:varchar(200);not_null"`
	TypeCategoryAbbr string `gorm:"column:typecategoryabbr;type:varchar(50);not_null"`
}

func (c TableTypeCategory) TableName() string {
	return "table_typecategory"
}
