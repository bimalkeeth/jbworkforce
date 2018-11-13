package entities

import "github.com/jinzhu/gorm"

type TablePayrollPayCode struct {
	gorm.Model
	PayrollPayCodeGroupId    uint                     `gorm:"column:payrollpaycodegroupid"`
	RosterTypeId             uint                     `gorm:"column:rostertypeid"`
	AllowanceId              uint                     `gorm:"column:allowanceid"`
	PayrollPayCodeAbbr       string                   `gorm:"column:payrollpaycodeabbr"`
	PayrollPayCodeName       string                   `gorm:"column:payrollpaycodename"`
	Description              string                   `gorm:"column:description"`
	IsActive                 bool                     `gorm:"column:isactive"`
	IsVerified               bool                     `gorm:"column:isverified"`
	UserId                   uint                     `gorm:"column:userid"`
	TablePayrollPayCodeGroup TablePayrollPayCodeGroup `gorm:"foreignkey:payrollpaycodegroupid"`
}

func (c TablePayrollPayCode) TableName() string {
	return "table_payrollpaycode"
}
