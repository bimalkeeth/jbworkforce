package entities

import "github.com/jinzhu/gorm"

type TableStatus struct {
	gorm.Model
	StatusName string `gorm:"column:statusname;not_null"`
	StatusAbbr string `gorm:"column:statusabbr;not_null"`
}

func (c TableStatus) TableName() string {
	return "table_status"
}
