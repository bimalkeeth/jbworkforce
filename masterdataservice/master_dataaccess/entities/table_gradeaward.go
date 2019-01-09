package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableGradeAward struct {
	gorm.Model
	GradeAwardName string    `gorm:"column:gradeawardname;type:varchar(200);not_null"`
	GradeAwardAbbr string    `gorm:"column:gradeawardabbr;type:varchar(50);not_null"`
	Description    string    `gorm:"column:description;type:varchar(300)"`
	ActiveFrom     time.Time `gorm:"column:activefrom;not_null"`
	ActiveTill     time.Time `gorm:"column:activetill;not_null"`
	DisplayOrder   int       `gorm:"column:displayorder;not_null"`
	UpdateUserId   uint      `gorm:"column:updateuserid;not_null"`
}

func (c TableGradeAward) TableName() string {
	return "table_gradeaward"
}
