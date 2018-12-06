package entities

import "github.com/jinzhu/gorm"

type TableClientSkillGroup struct {
	gorm.Model
	ClientSkillGroupName string                 `gorm:"column:clientskillgroupname;not_null"`
	ClientSkillTypes     []TableClientSkillType `gorm:"foreignkey:clientskillgroupid"`
}

func (c TableClientSkillGroup) TableName() string {
	return "table_clientskillgroup"
}
