package entities

import "github.com/jinzhu/gorm"

type TableClientSkillGroup struct {
	gorm.Model
	ClientSkillGroupName string `gorm:"column:clientskillgroupname;not_null"`
}

func (c TableClientSkillGroup) TableName() string {
	return "table_clientskillgroup"
}
