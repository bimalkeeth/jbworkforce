package entities

import "github.com/jinzhu/gorm"

type TableSmsGroup struct {
	gorm.Model
	SmsText      string `gorm:"column:smstext"`
	ShiftText    string `gorm:"column:shifttext"`
	SmsStatusId  uint   `gorm:"column:smsstatusid;not_null"`
	SmsGroupType string `gorm:"column:smsgrouptype;not_null"`
	Interval     int    `gorm:"column:interval;not_null"`
	CountDown    int    `gorm:"column:countdown;default:0"`
	SmsStatus    TableSmsStatus
}

func (c TableSmsGroup) TableName() string {
	return "table_smsgroup"
}
