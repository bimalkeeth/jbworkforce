package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableSmsConversation struct {
	gorm.Model
	SmsGroupId         uint      `gorm:"column:smsgroupid"`
	WorkforceRequestId uint      `gorm:"column:workforcerequestid;not_null"`
	EmployeeId         uint      `gorm:"column:employeeid;not_null"`
	DepartmentId       uint      `gorm:"column:departmentid;not_null"`
	ShiftTypeId        uint      `gorm:"column:shifttypeid;not_null"`
	DateWeekId         uint      `gorm:"column:dateweekid;not_null"`
	StartTime          time.Time `gorm:"column:starttime"`
	EndTime            time.Time `gorm:"column:endtime"`
	SmsStatusId        uint      `gorm:"column:smsstatusid;not_null"`
}

func (c TableSmsConversation) TableName() string {
	return "table_smsconversation"
}
