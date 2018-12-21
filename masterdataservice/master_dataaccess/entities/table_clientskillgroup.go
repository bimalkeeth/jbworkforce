package entities

import "github.com/jinzhu/gorm"

type TableClientSkillGroup struct {
	gorm.Model
	ClientSkillGroupName string `gorm:"column:clientskillgroupname;type:varchar(100);not_null"`
	ClientSkillTypes     []TableClientSkillType
}

func (c TableClientSkillGroup) TableName() string {
	return "table_clientskillgroup"
}
