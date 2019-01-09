package entities

import "github.com/jinzhu/gorm"

type TableGradeWorkforceGroup struct {
	gorm.Model
	GradeWorkforceGroupName string `gorm:"column:gradeworkforcegroupname;type:varchar(200);not_null"`
}

func (c TableGradeWorkforceGroup) TableName() string {
	return "table_gradeworkforcegroup"
}
