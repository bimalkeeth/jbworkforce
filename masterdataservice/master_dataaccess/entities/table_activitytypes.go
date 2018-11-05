package entities

import "github.com/jinzhu/gorm"

type TableActivityTypes struct {
	gorm.Model
	ActivityTypeName string `gorm:"column:activitytypename;not_null"`
	ActivityAbbr     string `gorm:"column:activityabbr;not_null"`
}

func (c TableActivityTypes) TableName() string {
	return "table_activitytypes"
}
