package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableSmsLog struct {
	gorm.Model
	MessageId    uint      `gorm:"column:messageid;not_null"`
	MobileNumber string    `gorm:"column:mobilenumber;not_null"`
	Message      string    `gorm:"column:message;not_null"`
	DateTime     time.Time `gorm:"column:datetime;not_null"`
	SmsLogTypeId uint      `gorm:"column:smslogtypeid;not_null"`
	SendStatus   string    `gorm:"column:sendstatus;not_null"`
	SmslogType   TableSmslogType
}

func (c TableSmsLog) TableName() string {
	return "table_smslog"
}
