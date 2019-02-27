package entities

import "github.com/jinzhu/gorm"

type TableHealthGroup struct {
	gorm.Model
	HealthGroupName string `gorm:"column:healthgroupname;type:varchar(200);not_null"`
	HealthGroupAbbr string `gorm:"column:healthgroupabbr;type:varchar(100);not_null"`
	Description     string `gorm:"column:description;not_null"`
	UserId          uint   `gorm:"column:userId;not_null"`
	ClinicalStreams []*TableClinicalStream
}

func (c TableHealthGroup) TableName() string {
	return "table_healthgroup"
}
