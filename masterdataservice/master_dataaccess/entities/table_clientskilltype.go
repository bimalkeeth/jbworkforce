package entities

import "github.com/jinzhu/gorm"

type TableClientSkillType struct {
	gorm.Model
	ClientSkillGroupId  uint   `gorm:"column:clientskillgroupid;not_null"`
	ClientSkillTypeName string `gorm:"column:clientskilltypename;type:varchar(100);not_null"`
	ClientSkillGroup    *TableClientSkillGroup
}

func (c TableClientSkillType) TableName() string {
	return "table_clientskilltype"
}
