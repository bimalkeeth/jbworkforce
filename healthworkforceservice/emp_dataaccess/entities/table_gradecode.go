package entities

import (
	"github.com/jinzhu/gorm"
)

type TableGradeCode struct {
	gorm.Model
	GradeProfessionId      uint                   `gorm:"column:gradeprofessionid;not_null"`
	GradeCodeName          string                 `gorm:"column:gradecodename;not_null"`
	GradeCodeAbbr          string                 `gorm:"column:gradecodeabbr;not_null"`
	Description            string                 `gorm:"column:description"`
	DisplayControl         string                 `gorm:"column:displaycontrol"`
	DisplayOrder           int                    `gorm:"column:displayorder;not_null"`
	GradeCodeSettingId     uint                   `gorm:"column:gradecodesettingid;not_null"`
	PayrollGradeCode       string                 `gorm:"column:payrollgradecode;not_null"`
	IsActive               bool                   `gorm:"column:isactive;not_null"`
	IsYpIncluded           bool                   `gorm:"column:isypincluded;not_null"`
	IsStaffIncluded        bool                   `gorm:"column:isstaffincluded;not_null"`
	TableGradeCodeSettings TableGradeCodeSettings `gorm:"foreignkey:gradecodesettingid"`
}

func (c TableGradeCode) TableName() string {
	return "table_gradecode"
}
