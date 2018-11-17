package entities

import "github.com/jinzhu/gorm"

type TableWeeklyData struct {
	gorm.Model
	ReferenceId            uint           `gorm:"column:referenceid"`
	EmployeeId             uint           `gorm:"column:employeeid;not_null"`
	SubDepartmentId        uint           `gorm:"column:subdepartmentid;not_null"`
	DateWeekId             uint           `gorm:"column:dateweekid;not_null"`
	RosterTypeId           uint           `gorm:"column:rostertypeid;not_null"`
	HealthCategoryId       uint           `gorm:"column:healthcategoryid;not_null"`
	SkillTypeId            uint           `gorm:"column:skilltypeid;not_null"`
	GradeCodeId            uint           `gorm:"column:gradecodeid;not_null"`
	EmployeeTypeId         uint           `gorm:"column:employeetypeid;not_null"`
	ClinicalStreamId       uint           `gorm:"column:clinicalstreamid;not_null"`
	Hours                  float64        `gorm:"column:hours;default:0.0"`
	DataType               string         `gorm:"column:datatype"`
	Eft                    float64        `gorm:"column:eft;not_null;default:0.0"`
	AnnualLeaveRequirement float64        `gorm:"column:annualleaverequirement;default:0.0"`
	IsVirtual              bool           `gorm:"column:IsVirtual;default:false"`
	IsProfiled             bool           `gorm:"column:isprofiled;default:false"`
	IsActive               bool           `gorm:"column:isactive;;default:true"`
	DataOption             int            `gorm:"column:dataoption;not_null"`
	TableCatStream         TableCatStream `gorm:"foreignkey:clinicalstreamid"`
}

func (c TableWeeklyData) TableName() string {
	return "table_weeklydata"
}
