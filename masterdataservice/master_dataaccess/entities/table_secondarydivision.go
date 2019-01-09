package entities

import "github.com/jinzhu/gorm"

type TableSecondaryDivision struct {
	gorm.Model
	SecondaryDivisionName string `gorm:"column:secondarydivisionname;type:varchar(200);not_null"`
	SecondaryDivisionAbbr string `gorm:"column:secondarydivisionabbr;type:varchar(50);not_null"`
	IsActive              bool   `gorm:"column:isactive;default:true"`
}

func (c TableSecondaryDivision) TableName() string {
	return "table_secondarydivision"
}
