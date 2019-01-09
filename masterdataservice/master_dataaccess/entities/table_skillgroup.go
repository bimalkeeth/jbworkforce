package entities

import "github.com/jinzhu/gorm"

type TableSkillGroup struct {
	gorm.Model
	SkillGroupName string `gorm:"column:skillgroupname;type:varchar(200);not_null"`
	Description    string `gorm:"column:description;type:varchar(400)"`
	DisplayOrder   int    `gorm:"column:displayorder;default:0"`
	UserId         uint   `gorm:"column:userid;not_null"`
	SkillTypes     []TableSkillType
}

func (c TableSkillGroup) TableName() string {
	return "table_skillgroup"
}
