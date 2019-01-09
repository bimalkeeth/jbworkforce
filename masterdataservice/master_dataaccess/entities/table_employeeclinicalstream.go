package entities

import "github.com/jinzhu/gorm"

type TableEmployeeClinicalStream struct {
	gorm.Model
	EmployeeId       uint `gorm:"column:employeeid;not_null"`
	ClinicalStreamId uint `gorm:"column:clinicalstreamid;not_null"`
	Employee         TableEmployee
}

func (c TableEmployeeClinicalStream) TableName() string {
	return "table_employeeclinicalstream"
}
