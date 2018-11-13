package entities

import "github.com/jinzhu/gorm"

type TablePayrollPayCodeGroup struct {
	gorm.Model
	PayrollPayCodeGroupName string `gorm:"column:payrollpaycodegroupname"`
}

func (c TablePayrollPayCodeGroup) TableName() string {
	return "table_payrollpaycodegroup"
}
