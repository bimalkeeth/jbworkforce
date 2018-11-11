package entities

import "github.com/jinzhu/gorm"

type TableGradeCodeAward struct {
	gorm.Model
	GradeCodeId  uint    `gorm:"column:gradecodeid;not_null"`
	GradeAwardId uint    `gorm:"column:gradeawardid;not_null"`
	HourlyRate   float64 `gorm:"column:hourlyrate;not_null"`
	UpdateUserId uint    `gorm:"column:updateuserid;not_null"`
}

func (c TableGradeCodeAward) TableName() string {
	return "table_gradecodeaward"
}
