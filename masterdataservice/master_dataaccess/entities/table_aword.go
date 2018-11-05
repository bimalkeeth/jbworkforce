package entities

import "github.com/jinzhu/gorm"

type TableAward struct {
	gorm.Model
	AwardName string `gorm:"column:awardname;not_null"`
	AwardAbbr string `gorm:"column:awardabbr;not_null"`
}

func (c TableAward) TableName() string {
	return "table_award"
}
