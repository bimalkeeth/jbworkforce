package entities

import "github.com/jinzhu/gorm"

type TableGradeAllowanceFactor struct {
	gorm.Model
	GradeAllowanceGroupId uint    `gorm:"column:gradeallowancegroupid;not_null"`
	GradeAwardId          uint    `gorm:"column:gradeawardid;not_null"`
	Factor                float64 `gorm:"column:factor;default:0.0"`
	Description           string  `gorm:"column:description"`
}

func (c TableGradeAllowanceFactor) TableName() string {
	return "table_gradeallowancefactor"
}
