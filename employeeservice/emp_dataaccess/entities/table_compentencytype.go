package entities

import "github.com/jinzhu/gorm"

type TableCompentencyType struct {
	gorm.Model
	CompentencyName string `gorm:"column:compentencyname;not_null"`
	WeekLength      int    `gorm:"column:weeklength;not_null"`
	CompentencyAbbr string `gorm:"column:compentencyabbr;not_null"`
}

func (c TableCompentencyType) TableName() string {
	return "table_compentencytype"
}
