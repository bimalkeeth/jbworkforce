package entities

import "github.com/jinzhu/gorm"

type TableRationaleType struct {
	gorm.Model
	RationaleGroupName string `gorm:"column:rationalegroupname"`
	RationaleTypeName  string `gorm:"column:rationaletypename;not_null"`
	RationaleTypeAbbr  string `gorm:"column:rationaletypeabbr;not_null"`
	Description        string `gorm:"column:description"`
	IsActive           bool   `gorm:"column:isactive;default:true"`
	IsSpecial          bool   `gorm:"column:isspecial;default:false"`
	DisplayOrder       int    `gorm:"column:displayorder;default:0"`
}

func (c TableRationaleType) TableName() string {
	return "table_rationaletype"
}
