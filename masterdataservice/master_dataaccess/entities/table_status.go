package entities

import "github.com/jinzhu/gorm"

type TableStatus struct {
	gorm.Model
	StatusName      string          `gorm:"column:statusname;not_null"`
	StatusAbbr      string          `gorm:"column:statusabbr;not_null"`
	StatusTypeId    uint            `gorm:"column:statustypeid;not_null"`
	TableStatusType TableStatusType `gorm:"foreignkey:statustypeid"`
}

func (c TableStatus) TableName() string {
	return "table_status"
}
