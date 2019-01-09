package entities

import "github.com/jinzhu/gorm"

type TableGradeProfession struct {
	gorm.Model
	GradeAllowanceGroupId      uint    `gorm:"column:gradeallowancegroupid;not_null"`
	GradeWorkforceGroupId      uint    `gorm:"column:gradeWorkforcegroupid;not_null"`
	GradeProfessionName        string  `gorm:"column:gradeprofessionname;type:varchar(250);not_null"`
	GradeProfessionAbbr        string  `gorm:"column:gradeprofessionabbr;type:varchar(50);not_null"`
	Description                string  `gorm:"column:description;type:varchar(300);"`
	DisplayOrder               int     `gorm:"column:displayorder;default:0"`
	FullTimeAnnualLeaveInWeeks float64 `gorm:"column:fulltimeannualleaveinWeeks;default:0.0"`
	PartTimeAnnualLeaveInWeeks float64 `gorm:"column:parttimeannualleaveinweeks;default:0.0"`
	AllowanceGroup             TableAllowanceGroup
	GradeWorkforceGroup        TableGradeWorkforceGroup
}

func (c TableGradeProfession) TableName() string {
	return "table_gradeprofession"
}
