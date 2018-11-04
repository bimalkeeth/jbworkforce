package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableEmployeeQualifications struct {
	gorm.Model
	QualificationId     uint                `gorm:"column:qualificationid;not_null"`
	EmployeeId          uint                `gorm:"column:employeeid;not_null"`
	RegistrationDate    time.Time           `gorm:"column:registrationdate"`
	ExpireDate          time.Time           `gorm:"column:expiredate"`
	Institution         string              `gorm:"column:institution;not_null"`
	InstitutionAddress  string              `gorm:"column:institutionaddress"`
	InstitutionContact  string              `gorm:"column:institutioncontact"`
	TableEmployee       TableEmployee       `gorm:"foreignkey:employeeid"`
	TableQualifications TableQualifications `gorm:"foreignkey:qualificationid"`
}

func (c TableEmployeeQualifications) TableName() string {
	return "table_employeequalifications"
}
