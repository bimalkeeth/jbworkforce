package entities

import "github.com/jinzhu/gorm"

type TablePayrollActivityRaw struct {
	gorm.Model
	EmployeeNumber       uint    `gorm:"column:employeenumber"`
	EmployeeName         string  `gorm:"column:employeename"`
	PayPeriodId          string  `gorm:"column:payperiodid"`
	PayPeriodSubId       uint    `gorm:"column:payperiodsubid"`
	ForPayperiodId       string  `gorm:"column:forpayperiodid"`
	ForPayPeriodSubId    uint    `gorm:"column:forpayperiodsubid"`
	MasterFileCostCentre string  `gorm:"column:masterfilecostcentre"`
	ActualCostCentre     string  `gorm:"column:actualcostcentre"`
	GradeCode            string  `gorm:"column:gradecode"`
	GradeSubCode         uint    `gorm:"column:gradesubcode"`
	PayType              string  `gorm:"column:paytype"`
	PayTypeDescription   string  `gorm:"column:paytypedescription"`
	PayTypeRate          float64 `gorm:"column:paytyperate"`
	PayTypeHour          float64 `gorm:"column:paytypehour"`
	AmountPaid           float64 `gorm:"column:amountpaid"`
}

func (c TablePayrollActivityRaw) TableName() string {
	return "table_payrollactivityraw"
}
