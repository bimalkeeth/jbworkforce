package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableWorkforceRosterShiftClassified struct {
	gorm.Model
	SkillGroupId         uint                 `gorm:"column:skillgroupid;not_null"`
	WorkforceRosterId    uint                 `gorm:"column:workforcerosterid;not_null"`
	StartTime            time.Time            `gorm:"column:starttime;not_null"`
	EndTime              time.Time            `gorm:"column:endtime;not_null"`
	RequiredEmployees    int                  `gorm:"column:requiredemployees;not_null"`
	TableWorkforceRoster TableWorkforceRoster `gorm:"foreignkey:workforcerosterid"`
}

func (c TableWorkforceRosterShiftClassified) TableName() string {
	return "table_workforcerostershiftclassified"
}
