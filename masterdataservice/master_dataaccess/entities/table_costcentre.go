package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableCostCentre struct {
	gorm.Model
	HospitalId       uint    `gorm:"column:hospitalid;not_null"`
	CostCentreName   string  `gorm:"column:costcentrename;type:varchar(200);not_null"`
	CostCentreAbbr   string  `gorm:"column:costcentreabbr;type:varchar(128);not_null"`
	CostCentreType   string  `gorm:"column:costcentretype;type:varchar(400)"`
	Description      string  `gorm:"column:description"`
	IsActive         bool    `gorm:"column:isactive;default:false"`
	InPayroll        bool    `gorm:"column:inpayroll;default:false"`
	RosterEft        float64 `gorm:"column:rostereft;default:0.00"`
	BudgetEft        float64 `gorm:"column:budgeteft;default:0.00"`
	BudgetBaseEft    float64 `gorm:"column:budgetbaseeft;default:0.00"`
	BudgetDollar     float64 `gorm:"column:budgetdollar;default:0.00"`
	TableHospital    *TableHospital
	SubDepartments   []*TableSubDepartment
	EmployeeEmpTypes []*TableEmployeeEmpType
}

func (c TableCostCentre) TableName() string {
	return "table_costcentre"
}

func (c TableCostCentre) Validate(db *gorm.DB) {
	if len(c.CostCentreAbbr) > 128 {
		_ = db.AddError(errors.New("costcentre abbr length should be less or equal to 128"))
	}
	if c.CostCentreAbbr == "" {
		_ = db.AddError(errors.New("costcentre abbr should not be empty"))
	}
	if len(c.CostCentreName) > 200 {
		_ = db.AddError(errors.New("costcentre name length should be less or equal to 200"))
	}
	if c.CostCentreName == "" {
		_ = db.AddError(errors.New("costcentre name should not be empty"))
	}
	if len(c.CostCentreType) > 400 {
		_ = db.AddError(errors.New("description length should be less or equal to 400"))
	}
}
