package entities

import "github.com/jinzhu/gorm"

type TableYearlyPlannerLeave struct {
	gorm.Model
	YearlyPlannerWeekId         uint                        `gorm:"column:yearlyplannerweekid;not_null"`
	YearlyPlannerLeaveFormId    uint                        `gorm:"column:yearlyplannerleaveformid;not_null"`
	SubDepartmentId             uint                        `gorm:"column:subdepartmentid;not_null"`
	DateWeekId                  uint                        `gorm:"column:dateweekid;not_null"`
	RosterTypeId                uint                        `gorm:"column:rostertypeid;not_null"`
	DailyLeaveHour              float64                     `gorm:"column:dailyleavehour;default:0.0"`
	TableYearlyPlannerWeek      TableYearlyPlannerWeek      `gorm:"foreignkey:yearlyplannerweekid"`
	TableYearlyPlannerLeaveForm TableYearlyPlannerLeaveForm `gorm:"foreignkey:yearlyplannerleaveformid"`
}

func (c TableYearlyPlannerLeave) TableName() string {
	return "table_yearlyplannerleave"
}
