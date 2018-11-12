package entities

import "github.com/jinzhu/gorm"

type TableEmployeeClinicalStream struct {
	gorm.Model
	EmployeeId           uint          `gorm:"column:employeeid;not_null"`
	ProfessionalStreamId uint          `gorm:"column:professionalstreamid;not_null"`
	TableEmployee        TableEmployee `gorm:"foreignkey:employeeid"`
}

func (c TableEmployeeClinicalStream) TableName() string {
	return "table_employeeclinicalstream"
}
