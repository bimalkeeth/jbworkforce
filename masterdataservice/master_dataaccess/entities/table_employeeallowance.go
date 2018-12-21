package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableEmployeeAllowance struct {
	gorm.Model
	EmployeeId  uint      `gorm:"column:employeeid;not_null"`
	AllowanceId uint      `gorm:"column:allowanceid;not_null"`
	StartDate   time.Time `gorm:"column:startdate;not_null"`
	EndDate     time.Time `gorm:"column:enddate"`
	Dollar      float64   `gorm:"column:dollar;not_null;default:0.00"`
	IsActive    bool      `gorm:"column:isactive;default:true"`
	Employee    TableEmployee
	Allowance   TableAllowance
}

func (c TableEmployeeAllowance) TableName() string {
	return "table_employeeallowance"
}
