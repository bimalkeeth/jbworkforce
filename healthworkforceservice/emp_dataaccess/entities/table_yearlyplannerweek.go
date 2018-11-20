package entities

import "github.com/jinzhu/gorm"

type TableYearlyPlannerWeek struct {
	gorm.Model
	EmployeeId   uint `gorm:"column:employeeid;not_null"`
	DateWeekId   uint `gorm:"column:dateweekid;not_null"`
	IsAutoSync   bool `gorm:"column:isautosync;default:true"`
	IsUserChange bool `gorm:"column:isuserchange;default:false"`
	UserId       uint `gorm:"column:userid;not_null"`
}

func (c TableYearlyPlannerWeek) TableName() string {
	return "table_yearlyplannerweek"
}
