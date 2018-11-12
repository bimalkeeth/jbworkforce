package entities

import "github.com/jinzhu/gorm"

type TableGradeWorkforceAllowance struct {
	gorm.Model
	GradeWorkforceGroupId uint    `gorm:"column:gradeworkforcegroupid;not_null"`
	StaffTypeId           uint    `gorm:"column:stafftypeid;not_null"`
	AllowanceId           uint    `gorm:"column:allowanceid;not_null"`
	GradeAwardId          uint    `gorm:"column:gradeawardid;not_null"`
	RateCalculatorId      uint    `gorm:"column:ratecalculatorid;not_null"`
	Rate                  float64 `gorm:"column:rate;default:0.0"`
	Description           string  `gorm:"column:description"`
}

func (c TableGradeWorkforceAllowance) TableName() string {
	return "table_gradeworkforceallowance"
}
