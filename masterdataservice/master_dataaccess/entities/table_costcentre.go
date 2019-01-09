package entities

import "github.com/jinzhu/gorm"

type TableCostCentre struct {
	gorm.Model
	HospitalId       uint    `gorm:"column:hospitalid;not_null"`
	CostCentreName   string  `gorm:"column:costcentrename;type:varchar(200);not_null"`
	CostCentreAbbr   string  `gorm:"column:costcentreabbr;type:varchar(50);not_null"`
	CostCentreType   string  `gorm:"column:costcentretype;type:varchar(400)"`
	Description      string  `gorm:"column:description"`
	IsActive         bool    `gorm:"column:isactive;default:false"`
	InPayroll        bool    `gorm:"column:inpayroll;default:false"`
	RosterEft        float64 `gorm:"column:rostereft;default:0.00"`
	BudgetEft        float64 `gorm:"column:budgeteft;default:0.00"`
	BudgetBaseEft    float64 `gorm:"column:budgetbaseeft;default:0.00"`
	BudgetDollar     float64 `gorm:"column:budgetdollar;default:0.00"`
	TableHospital    TableHospital
	SubDepartments   []TableSubDepartment
	EmployeeEmpTypes []TableEmployeeEmpType
}

func (c TableCostCentre) TableName() string {
	return "table_costcentre"
}
