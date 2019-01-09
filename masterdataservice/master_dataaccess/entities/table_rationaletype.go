package entities

import "github.com/jinzhu/gorm"

type TableRationaleType struct {
	gorm.Model
	RationaleGroupName string `gorm:"column:rationalegroupname;type:varchar(150);"`
	RationaleTypeName  string `gorm:"column:rationaletypename;type:varchar(150);not_null"`
	RationaleTypeAbbr  string `gorm:"column:rationaletypeabbr;type:varchar(50);not_null"`
	Description        string `gorm:"column:description;type:varchar(300);"`
	IsActive           bool   `gorm:"column:isactive;default:true"`
	IsSpecial          bool   `gorm:"column:isspecial;default:false"`
	DisplayOrder       int    `gorm:"column:displayorder;default:0"`
}

func (c TableRationaleType) TableName() string {
	return "table_rationaletype"
}
