package entities

import "github.com/jinzhu/gorm"

type TableEmployeeAddress struct {
	gorm.Model
	AddressId     uint          `gorm:"column:addressid;not_null"`
	EmployeeId    uint          `gorm:"column:employeeid;not_null"`
	TableAddress  TableAddress  `gorm:"foreignkey:addressid"`
	TableEmployee TableEmployee `gorm:"foreignkey:employeeid"`
}

func (c TableEmployeeAddress) TableName() string {
	return "table_employeeaddress"
}
