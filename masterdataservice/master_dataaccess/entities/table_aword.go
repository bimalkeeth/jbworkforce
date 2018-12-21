package entities

import "github.com/jinzhu/gorm"

type TableAward struct {
	gorm.Model
	AwardName string `gorm:"column:awardname;type:varchar(100);not_null"`
	AwardAbbr string `gorm:"column:awardabbr;type:varchar(100);not_null"`
}

func (c TableAward) TableName() string {
	return "table_award"
}
