package entities

import "github.com/jinzhu/gorm"

type TableWorkforceAvailability struct {
	gorm.Model
	WorkforceRequestId      uint                  `gorm:"column:workforcerequestid;not_null"`
	EmployeeId              uint                  `gorm:"column:employeeid;not_null"`
	DateWeekId              uint                  `gorm:"column:dateweekid;not_null"`
	AvailabilityStatusId    uint                  `gorm:"column:availabilitystatusid;not_null"`
	OrgAvailabilityStatusId uint                  `gorm:"column:templatespreadname;not_null"`
	EmployeeTypeId          uint                  `gorm:"column:employeetypeid;not_null"`
	IsPreBooked             bool                  `gorm:"column:isprebooked;default:false"`
	MigrationId             int                   `gorm:"column:migrationid;default:0"`
	BookedByMemberId        uint                  `gorm:"column:bookedbymemberid"`
	Note                    string                `gorm:"column:note"`
	TableWorkforceRequest   TableWorkforceRequest `gorm:"foreignkey:workforcerequestid"`
}

func (c TableWorkforceAvailability) TableName() string {
	return "table_workforceavailability"
}
