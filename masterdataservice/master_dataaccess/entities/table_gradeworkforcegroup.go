package entities

import "github.com/jinzhu/gorm"

type TableGradeWorkforceGroup struct {
	gorm.Model
	GradeWorkforceGroupName string `gorm:"column:gradeworkforcegroupname;not_null"`
}

func (c TableGradeWorkforceGroup) TableName() string {
	return "table_gradeworkforcegroup"
}
