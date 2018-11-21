package entities

import "github.com/jinzhu/gorm"

type TableGradeAllowanceAward struct {
	gorm.Model
	GradeAllowanceGroupId uint    `gorm:"column:gradeallowancegroupid;not_null"`
	AllowanceId           uint    `gorm:"column:allowanceid;not_null"`
	GradeAwardId          uint    `gorm:"column:gradeawardid;not_null"`
	AllowanceLevelId      uint    `gorm:"column:allowancelevelid;not_null"`
	RateCalculatorId      uint    `gorm:"column:ratecalculatorid;not_null"`
	Rate                  float64 `gorm:"column:rate;not_null"`
	Description           string  `gorm:"column:description;not_null"`
}

func (c TableGradeAllowanceAward) TableName() string {
	return "table_gradeallowanceaward"
}
