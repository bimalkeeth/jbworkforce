package entities

import "github.com/jinzhu/gorm"

type TableYearlyPlannerRoster struct {
	gorm.Model
	YearlyPlannerWeekId    uint                   `gorm:"column:yearlyplannerweekid;not_null"`
	RosterTypeId           uint                   `gorm:"column:rostertypeid;not_null"`
	SubDepartmentId        uint                   `gorm:"column:subdepartmentid;not_null"`
	GradeCodeId            uint                   `gorm:"column:gradecodeid;not_null"`
	EmployeeTypeId         uint                   `gorm:"column:employeetypeid;not_null"`
	HealthCategoryId       uint                   `gorm:"column:healthcategoryid;not_null"`
	ClinicalStreamId       uint                   `gorm:"column:clinicalstreamid;not_null"`
	SkillTypeId            uint                   `gorm:"column:skilltypeid;not_null"`
	ActualWeeklyHour       float64                `gorm:"column:actualweeklyhour;default:0.0"`
	Version                int                    `gorm:"column:version;default:0"`
	TableGradeCode         TableGradeCode         `gorm:"foreignkey:gradecodeid"`
	TableYearlyPlannerWeek TableYearlyPlannerWeek `gorm:"foreignkey:yearlyplannerweekid"`
}

func (c TableYearlyPlannerRoster) TableName() string {
	return "table_yearlyplannerroster"
}
