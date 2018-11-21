package entities

import "github.com/jinzhu/gorm"

type TableWorkforceRequestSkillType struct {
	gorm.Model
	WorkforceRequestId    uint                  `gorm:"column:workforcerequestid;not_null"`
	SkillTypeId           uint                  `gorm:"column:skilltypeid;not_null"`
	TableWorkforceRequest TableWorkforceRequest `gorm:"foreignkey:workforcerequestid"`
}

func (c TableWorkforceRequestSkillType) TableName() string {
	return "table_workforcerequestskilltype"
}
