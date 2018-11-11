package entities

import "github.com/jinzhu/gorm"

type TableGradeCodeSettings struct {
	gorm.Model
	ClinicalStreamId uint `gorm:"column:clinicalstreamid;not_null"`
	HealthCategoryId uint `gorm:"column:healthcategoryid;not_null"`
	SkillTypeId      uint `gorm:"column:skilltypeid;not_null"`
}

func (c TableGradeCodeSettings) TableName() string {
	return "table_gradecodesettings"
}
