package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableEmployeeDetail struct {
	gorm.Model
	EmployeeId             uint          `gorm:"column:employeeid;not_null;"`
	BestTimeToPhone        time.Time     `gorm:"column:besttimetophone"`
	PoliceCheckDate        time.Time     `gorm:"column:policecheckdate"`
	PoliceCheckComplete    bool          `gorm:"column:policecheckcomplete"`
	VisaIssuedDate         time.Time     `gorm:"column:visaissueddate"`
	RegistrationNumber     string        `gorm:"column:registrationnumber"`
	CanWorkNight           bool          `gorm:"column:canworknight;not_null;default:false"`
	CanWorkWeekEnd         bool          `gorm:"column:canworkweekend;not_null;default:false"`
	CanWorkSchoolHolidays  bool          `gorm:"column:canworkschoolholidays;not_null;default:true"`
	RegistrationIssuedDate time.Time     `gorm:"column:registrationissueddate"`
	RegistrationExpireDate time.Time     `gorm:"column:registrationexpiredate"`
	NoLift                 bool          `gorm:"column:nolift;not_null;default:true"`
	IsCitizen              bool          `gorm:"column:iscitizen;not_null;default:true"`
	TableEmployee          TableEmployee `gorm:"foreignkey:employeeid"`
}

func (c TableEmployeeDetail) TableName() string {
	return "table_employeedetail"
}
