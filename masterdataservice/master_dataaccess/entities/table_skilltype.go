package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type TableSkillType struct {
	gorm.Model
	SkillGroupId       uint           `gorm:"column:skillgroupid;not_null"`
	ClientSkillTypeId  uint           `gorm:"column:clientskilltypeid"`
	SkillTypeName      string         `gorm:"column:skilltypename;type:varchar(200);not_null"`
	SkillTypeAbbr      string         `gorm:"column:skilltypeabbr;type:varchar(50);not_null"`
	Description        string         `gorm:"column:description;type:varchar(400)"`
	DisplayOrder       int            `gorm:"column:displayorder;default:0"`
	Setting            postgres.Jsonb `gorm:"column:setting"`
	IsYpIncluded       bool           `gorm:"column:isypincluded;default:true"`
	SkillGroup         TableSkillGroup
	ClientSkillType    TableClientSkillType
	EmployeeSkillTypes []TableEmployeeSkillType
}

func (c TableSkillType) TableName() string {
	return "table_skilltype"
}
