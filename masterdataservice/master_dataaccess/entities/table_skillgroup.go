package entities

import "github.com/jinzhu/gorm"

type TableSkillGroup struct {
	gorm.Model
	SkillGroupName string `gorm:"column:skillgroupname;not_null"`
	Description    string `gorm:"column:description"`
	DisplayOrder   int    `gorm:"column:displayorder;default:0"`
	UserId         uint   `gorm:"column:userid;not_null"`
}

func (c TableSkillGroup) TableName() string {
	return "table_skillgroup"
}
