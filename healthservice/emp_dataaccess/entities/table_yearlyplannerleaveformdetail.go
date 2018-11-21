package entities

import "github.com/jinzhu/gorm"

type TableYearlyPlannerLeaveFormDetail struct {
	gorm.Model
	YearlyPlannerLeaveFormId    uint                        `gorm:"column:yearlyplannerleaveformid;not_null"`
	DateWeekId                  uint                        `gorm:"column:dateweekid;not_null"`
	RosterTypeId                uint                        `gorm:"column:rostertypeid;not_null"`
	Hour                        float64                     `gorm:"column:hour;default:0.0"`
	Description                 string                      `gorm:"column:description"`
	WeekLeaveType               string                      `gorm:"column:weekleavetype"`
	TableYearlyPlannerLeaveForm TableYearlyPlannerLeaveForm `gorm:"foreignkey:yearlyplannerleaveformid"`
}

func (c TableYearlyPlannerLeaveFormDetail) TableName() string {
	return "table_yearlyplannerleaveformdetail"
}
