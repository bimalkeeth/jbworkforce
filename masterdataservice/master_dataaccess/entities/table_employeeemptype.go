package entities

import "github.com/jinzhu/gorm"

type TableEmployeeEmpType struct {
	gorm.Model
	EmployeeId      uint `gorm:"column:employeeid;not_null"`
	EmployeeTypeId  uint `gorm:"column:employeetypeid;not_null"`
	SubDepartmentId uint `gorm:"column:subdepartmentid;not_null"`
	CostCenterId    uint `gorm:"column:costcenterid;not_null"`
	Employee        *TableEmployee
	EmployeeType    *TableEmployeeType
	SubDepartment   *TableSubDepartment
	CostCentre      *TableCostCentre
}

func (c TableEmployeeEmpType) TableName() string {
	return "table_employeeemptype"
}
