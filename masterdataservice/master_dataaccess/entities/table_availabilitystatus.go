package entities

import "github.com/jinzhu/gorm"

type TableAvailabilityStatus struct {
	gorm.Model
	AvailabilityStatusName string `gorm:"column:availabilitystatusname;type:varchar(100);not_null"`
}

func (c TableAvailabilityStatus) TableName() string {
	return "table_availabilitystatus"
}
