package entities

import "github.com/jinzhu/gorm"

type TableYearlyPlannerTemplate struct {
	gorm.Model
	YearlyPlannerWeekId    uint                   `gorm:"column:yearlyplannerweekid;not_null"`
	SubDepartmentId        uint                   `gorm:"column:subdepartmentid;not_null"`
	GradeCodeId            uint                   `gorm:"column:gradecodeid;not_null"`
	EmployeeTypeId         uint                   `gorm:"column:employeetypeid;not_null"`
	ContractedWeeklyHour   float64                `gorm:"column:contractedweeklyhour;default:0.0"`
	UserId                 uint                   `gorm:"column:userid;not_null"`
	TableYearlyPlannerWeek TableYearlyPlannerWeek `gorm:"foreignkey:yearlyplannerweekid"`
	TableGradeCode         TableGradeCode         `gorm:"foreignkey:gradecodeid"`
}

func (c TableYearlyPlannerTemplate) TableName() string {
	return "table_yearlyplannertemplate"
}
