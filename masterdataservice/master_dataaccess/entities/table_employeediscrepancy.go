package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableEmployeeDiscrepancy struct {
	gorm.Model
	EmployeeId        uint      `gorm:"column:employeeid;not_null"`
	DepartmentId      uint      `gorm:"column:departmentid;not_null"`
	OldValue          string    `gorm:"column:oldvalue"`
	NewValue          string    `gorm:"column:newvalue"`
	StatusId          uint      `gorm:"column:statusid;not_null"`
	Date              time.Time `gorm:"column:date"`
	DateModified      time.Time `gorm:"column:datemodified"`
	DiscrepancyTypeId uint      `gorm:"column:discrepancytypeid;not_null"`
	NewValueId        string    `gorm:"column:newvalueid;default:-1"`
	NewData           string    `gorm:"column:newdata"`
	Quality           string    `gorm:"column:quality"`
	IsReported        bool      `gorm:"column:isreported;default:false"`
	Status            *TableStatus
	Types             *TableTypes
	Employee          *TableEmployee
}

func (c TableEmployeeDiscrepancy) TableName() string {
	return "table_employeediscrepancy"
}
