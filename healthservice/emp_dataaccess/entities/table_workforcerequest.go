package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableWorkforceRequest struct {
	gorm.Model
	WorkforceRosterId          uint                       `gorm:"column:workforcerosterid;not_null"`
	WorkforceAvailabilityId    uint                       `gorm:"column:workforceavailabilityid"`
	DepartmentId               uint                       `gorm:"column:departmentid;not_null"`
	DateWeekId                 uint                       `gorm:"column:dateweekid;not_null"`
	ShiftTypeId                uint                       `gorm:"column:shifttypeid;not_null"`
	RationaleTypeId            uint                       `gorm:"column:rationaletypeid"`
	AllocationStatusId         uint                       `gorm:"column:allocationstatusid"`
	WorkforceEmployeeGroupId   uint                       `gorm:"column:workforceemployeegroupid"`
	GradeCodeId                uint                       `gorm:"column:gradecodeid"`
	EmployeeTypeId             uint                       `gorm:"column:employeetypeid"`
	StartTime                  time.Time                  `gorm:"column:starttime"`
	EndTime                    time.Time                  `gorm:"column:endtime"`
	IsMealBreak                bool                       `gorm:"column:ismealbreak;default:false"`
	ShiftLength                float64                    `gorm:"column:shiftlength;default:0.0"`
	Comment                    string                     `gorm:"column:comment"`
	IsCertificated             bool                       `gorm:"column:iscertificated;not_null"`
	IsRequired                 bool                       `gorm:"column:isrequired;not_null"`
	IsQualified                bool                       `gorm:"column:isqualified;not_null"`
	SuggestedStaffName         string                     `gorm:"column:suggestedstaffname"`
	SuggestedEmployeeId        uint                       `gorm:"column:suggestedemployeeid"`
	MigrationId                uint                       `gorm:"column:migrationid"`
	BookDate                   time.Time                  `gorm:"column:bookdate"`
	RequestDate                time.Time                  `gorm:"column:requestdate"`
	ConfirmationDate           time.Time                  `gorm:"column:confirmationdate"`
	ReCalc                     bool                       `gorm:"column:recalc"`
	TableWorkforceAvailability TableWorkforceAvailability `gorm:"foreignkey:workforceavailabilityid"`
	TableGradeCode             TableGradeCode             `gorm:"foreignkey:gradecodeid"`
}

func (c TableWorkforceRequest) TableName() string {
	return "table_workforcerequest"
}
