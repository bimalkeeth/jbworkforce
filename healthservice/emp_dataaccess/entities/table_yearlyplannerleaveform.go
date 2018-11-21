package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableYearlyPlannerLeaveForm struct {
	gorm.Model
	LeaveFormTitle     string    `gorm:"column:leaveformtitle;not_null"`
	LeaveFormName      string    `gorm:"column:leaveformname;not_null"`
	EmployeeId         uint      `gorm:"column:employeeid;not_null"`
	DepartmentId       uint      `gorm:"column:departmentid;not_null"`
	FirstLeaveDate     time.Time `gorm:"column:firstleavedate"`
	LastLeaveDate      time.Time `gorm:"column:lastleavedate"`
	FormSubmissionDate time.Time `gorm:"column:formsubmissiondate"`
	Description        string    `gorm:"column:description"`
}

func (c TableYearlyPlannerLeaveForm) TableName() string {
	return "table_yearlyplannerleaveform"
}
