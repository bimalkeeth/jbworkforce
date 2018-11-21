package entities

import "github.com/jinzhu/gorm"

type TableEmployeeGradeCode struct {
	gorm.Model
	EmployeeId  uint `gorm:"column:employeeid;not_null"`
	GradeCodeId uint `gorm:"column:gradecodeid;not_null"`
}

func (c TableEmployeeGradeCode) TableName() string {
	return "table_employeegradecode"
}
