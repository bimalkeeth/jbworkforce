package entities

import "github.com/jinzhu/gorm"

type TableEmployeeAgency struct {
	gorm.Model
	EmployeeId    uint          `gorm:"column:employeeid;not_null"`
	AgencyId      uint          `gorm:"column:agencyid;not_null"`
	TableEmployee TableEmployee `gorm:"foreignkey:employeeid"`
	TableAgency   TableAgency   `gorm:"foreignkey:agencyid"`
}

func (c TableEmployeeAgency) TableName() string {
	return "table_employeeagency"
}
