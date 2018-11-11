package entities

import "github.com/jinzhu/gorm"

type TableGradeAllowanceSpread struct {
	gorm.Model
	GradeAllowanceGroupId uint    `gorm:"column:gradeallowancegroupid;not_null"`
	GradeAwardId          uint    `gorm:"column:gradeawardid;not_null"`
	DateWeekId            int     `gorm:"column:dateweekid;not_null"`
	Factor                float64 `gorm:"column:factor;not_null"`
	Description           string  `gorm:"column:description"`
}

func (c TableGradeAllowanceSpread) TableName() string {
	return "table_gradeallowancespread"
}
