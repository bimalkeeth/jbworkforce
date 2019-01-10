package entities

import "github.com/jinzhu/gorm"

type TableStatus struct {
	gorm.Model
	StatusName   string `gorm:"column:statusname;type:varchar(200);not_null"`
	StatusAbbr   string `gorm:"column:statusabbr;type:varchar(50);not_null"`
	StatusTypeId uint   `gorm:"column:statustypeid;not_null"`
	StatusType   TableStatusType
}

func (c TableStatus) TableName() string {
	return "table_status"
}
