package entities

import "github.com/jinzhu/gorm"

type TableTypes struct {
	gorm.Model
	TypeName       string `gorm:"column:typename;type:varchar(200);not_null"`
	TypeAbbr       string `gorm:"column:typeabbr;type:varchar(50);not_null"`
	TypeCategoryId uint   `gorm:"column:typecategoryid;not_null"`
	TypeCategory   TableTypeCategory
}

func (c TableTypes) TableName() string {
	return "table_types"
}
