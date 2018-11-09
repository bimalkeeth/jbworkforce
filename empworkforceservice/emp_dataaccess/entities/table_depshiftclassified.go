package entities

import "github.com/jinzhu/gorm"

type TableDepShiftClassified struct {
	gorm.Model
	DepShiftId           uint                 `gorm:"column:depshiftid;not_null"`
	SkillGroupId         uint                 `gorm:"column:skillgroupid;not_null"`
	RequiredStaff        int                  `gorm:"column:requiredstaff;not_null;default:0"`
	TableDepartmentShift TableDepartmentShift `gorm:"foreignkey:depshiftid"`
}

func (c TableDepShiftClassified) TableName() string {
	return "table_depshiftclassified"
}
