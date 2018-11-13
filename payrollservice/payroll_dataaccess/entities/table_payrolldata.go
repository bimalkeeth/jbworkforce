package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TablePayrollData struct {
	gorm.Model
	PayPeriodNo      float64   `gorm:"column:payperiodno"`
	PayPeriodDate    time.Time `gorm:"column:payperioddate"`
	EmpNo            string    `gorm:"column:empno"`
	EmpName          string    `gorm:"column:empname"`
	AppointType      string    `gorm:"column:appointtype"`
	EFT              float64   `gorm:"column:eft"`
	GradeCode        string    `gorm:"column:gradecode"`
	GradeDescription string    `gorm:"column:gradedescription"`
	UnitId           string    `gorm:"column:unitid"`
	UnitName         string    `gorm:"column:unitname"`
	PayClass         string    `gorm:"column:payclass"`
	PayCode          string    `gorm:"column:paycode"`
	PayDescription   string    `gorm:"column:paydescription"`
	WorkedHrs        float64   `gorm:"column:workedhrs"`
	PayedAmount      float64   `gorm:"column:payedamount"`
	Tag              string    `gorm:"column:tag"`
}

func (c TablePayrollData) TableName() string {
	return "table_payrolldata"
}
