package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type TableSkillType struct {
	gorm.Model
	SkillGroupId         uint                     `gorm:"column:skillgroupid;not_null"`
	ClientSkillTypeId    uint                     `gorm:"column:clientskilltypeid"`
	SkillTypeName        string                   `gorm:"column:skilltypename;not_null"`
	SkillTypeAbbr        string                   `gorm:"column:skilltypeabbr;not_null"`
	Description          string                   `gorm:"column:description"`
	DisplayOrder         int                      `gorm:"column:displayorder;default:0"`
	Setting              postgres.Jsonb           `gorm:"column:setting"`
	IsYpIncluded         bool                     `gorm:"column:isypincluded;default:true"`
	TableSkillGroup      TableSkillGroup          `gorm:"foreignkey:skillgroupid"`
	TableClientSkillType TableClientSkillType     `gorm:"foreignkey:clientskilltypeid"`
	EmployeeSkillTypes   []TableEmployeeSkillType `gorm:"foreignkey:skilltypeid"`
}

func (c TableSkillType) TableName() string {
	return "table_skilltype"
}
