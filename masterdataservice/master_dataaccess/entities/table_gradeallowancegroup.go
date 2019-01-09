package entities

import "github.com/jinzhu/gorm"

type TableGradeAllowanceGroup struct {
	gorm.Model
	GradeAllowanceGroupName string `gorm:"column:gradeallowancegroupname;type:varchar(200);not_null"`
}

func (c TableGradeAllowanceGroup) TableName() string {
	return "table_gradeallowancegroup"
}
