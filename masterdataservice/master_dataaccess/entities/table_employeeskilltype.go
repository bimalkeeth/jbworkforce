package entities

import "github.com/jinzhu/gorm"

type TableEmployeeSkillType struct {
	gorm.Model
	EmployeeId     uint           `gorm:"column:employeeid;not_null"`
	SkillTypeId    uint           `gorm:"column:skilltypeid;not_null"`
	TableEmployee  TableEmployee  `gorm:"foreignkey:employeeid"`
	TableSkillType TableSkillType `gorm:"foreignkey:skilltypeid"`
}

func (c TableEmployeeSkillType) TableName() string {
	return "table_employeeskilltype"
}
