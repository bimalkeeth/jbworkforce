package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableSpecialRequest struct {
	gorm.Model
	WorkforceRequestId uint      `gorm:"column:workforcerequestid;not_null"`
	RequestedEmployee  string    `gorm:"column:requestedemployee"`
	AllocatedEmployee  string    `gorm:"column:allocatedemployee"`
	DateTime           time.Time `gorm:"column:datetime"`
}

func (c TableSpecialRequest) TableName() string {
	return "table_specialrequest"
}
