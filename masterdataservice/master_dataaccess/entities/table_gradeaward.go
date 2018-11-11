package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableGradeAward struct {
	gorm.Model
	GradeAwardName string    `gorm:"column:gradeawardname;not_null"`
	GradeAwardAbbr string    `gorm:"column:gradeawardabbr;not_null"`
	Description    string    `gorm:"column:description"`
	ActiveFrom     time.Time `gorm:"column:activefrom;not_null"`
	ActiveTill     time.Time `gorm:"column:activetill;not_null"`
	DisplayOrder   int       `gorm:"column:displayorder;not_null"`
	UpdateUserId   uint      `gorm:"column:updateuserid;not_null"`
}

func (c TableGradeAward) TableName() string {
	return "table_gradeaward"
}
