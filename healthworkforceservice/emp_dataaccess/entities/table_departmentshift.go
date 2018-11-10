package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableDepartmentShift struct {
	gorm.Model
	DepartmentId  uint      `gorm:"column:departmentid;not_null"`
	WeekdayId     uint      `gorm:"column:weekdayid;not_null"`
	ShiftTypeId   uint      `gorm:"column:shifttypeid;not_null"`
	RequiredStaff int       `gorm:"column:requiredstaff;not_null"`
	RosterStaff   int       `gorm:"column:rosterstaff;not_null"`
	StartTime     time.Time `gorm:"column:starttime"`
	EndTime       time.Time `gorm:"column:endtime"`
	StartDate     time.Time `gorm:"column:startdate"`
	EndDate       time.Time `gorm:"column:enddate"`
	ShiftHour     float64   `gorm:"column:shifthour"`
	IsActive      bool      `gorm:"column:isactive;default:false"`
}

func (c TableDepartmentShift) TableName() string {
	return "table_departshift"
}
