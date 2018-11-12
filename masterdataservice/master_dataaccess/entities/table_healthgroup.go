package entities

import "github.com/jinzhu/gorm"

type TableHealthGroup struct {
	gorm.Model
	HealthGroupName string `gorm:"column:healthgroupname;not_null"`
	HealthGroupAbbr string `gorm:"column:healthgroupabbr;not_null"`
	Description     string `gorm:"column:description;not_null"`
	UserId          uint   `gorm:"column:userId;not_null"`
}

func (c TableHealthGroup) TableName() string {
	return "table_healthgroup"
}
