package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type TableYearlyPlannerPayroll struct {
	gorm.Model
	EmployeeId       uint           `gorm:"column:employeeid;not_null"`
	DateWeekId       uint           `gorm:"column:dateweekid;not_null"`
	SubDepartmentId  uint           `gorm:"column:subdepartmentid;not_null"`
	RosterTypeId     uint           `gorm:"column:rostertypeid;not_null"`
	GradeCodeId      uint           `gorm:"column:gradecodeid;not_null"`
	EmployeeTypeId   uint           `gorm:"column:employeetypeid;not_null"`
	HealthCategoryId uint           `gorm:"column:healthcategoryid;not_null"`
	ClinicalStreamId uint           `gorm:"column:clinicalstreamid;not_null"`
	SkillTypeId      uint           `gorm:"column:skilltypeid;not_null"`
	Hour             float64        `gorm:"column:hour;default:0.0"`
	Comment          string         `gorm:"column:comment"`
	Setting          postgres.Jsonb `gorm:"column:setting"`
	TableGradeCode   TableGradeCode `gorm:"foreignkey:gradecodeid"`
}

func (c TableYearlyPlannerPayroll) TableName() string {
	return "table_yearlyplannerpayroll"
}
