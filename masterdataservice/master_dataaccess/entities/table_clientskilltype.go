package entities

import "github.com/jinzhu/gorm"

type TableClientSkillType struct {
	gorm.Model
	ClientSkillGroupId    uint                  `gorm:"column:clientskillgroupid;not_null"`
	ClientSkillTypeName   string                `gorm:"column:clientskilltypename;not_null"`
	TableClientSkillGroup TableClientSkillGroup `gorm:"foreignkey:clientskillgroupid"`
}

func (c TableClientSkillType) TableName() string {
	return "table_clientskilltype"
}
