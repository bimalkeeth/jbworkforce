package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableEmployeeQualifications struct {
	gorm.Model
	QualificationId    uint      `gorm:"column:qualificationid;not_null"`
	EmployeeId         uint      `gorm:"column:employeeid;not_null"`
	RegistrationDate   time.Time `gorm:"column:registrationdate"`
	ExpireDate         time.Time `gorm:"column:expiredate"`
	Institution        string    `gorm:"column:institution;type:varchar(250);not_null"`
	InstitutionAddress string    `gorm:"column:institutionaddress;type:varchar(300)"`
	InstitutionContact string    `gorm:"column:institutioncontact;type:varchar(200)"`
	Employee           *TableEmployee
}

func (c TableEmployeeQualifications) TableName() string {
	return "table_employeequalifications"
}
