package entities

import "github.com/jinzhu/gorm"

type TableWorkforceEmployeeGroup struct {
	gorm.Model
	WorkforceStaffGroupName string `gorm:"column:workforcestaffgroupname;not_null"`
}

func (c TableWorkforceEmployeeGroup) TableName() string {
	return "table_workforceemployeegroup"
}
