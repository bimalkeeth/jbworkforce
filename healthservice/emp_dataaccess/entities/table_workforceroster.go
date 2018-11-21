package entities

import "github.com/jinzhu/gorm"

type TableWorkforceRoster struct {
	gorm.Model
	DateWeekId    uint `gorm:"column:dateweekid;not_null"`
	ShiftTypeId   uint `gorm:"column:shifttypeid;not_null"`
	DepartmentId  uint `gorm:"column:departmentid;not_null"`
	BudgetedStaff int  `gorm:"column:budgetedstaff;default:0"`
	RequiredStaff int  `gorm:"column:requiredstaff;default:0"`
	RosteredStaff int  `gorm:"column:rosteredstaff;default:0"`
	MigrationId   uint `gorm:"column:MigrationId"`
}

func (c TableWorkforceRoster) TableName() string {
	return "table_workforceroster"
}
