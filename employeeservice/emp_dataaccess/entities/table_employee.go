package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableEmployee struct {
	gorm.Model
	AddressId       uint      `gorm:"column:addressid;not_null"`
	FirstName       string    `gorm:"column:firstname;not_null"`
	LastName        string    `gorm:"column:lastname;not_null"`
	PrefName        string    `gorm:"column:name;not_null"`
	GenderId        uint      `gorm:"column:genderid;not_null"`
	DateOfBirth     time.Time `gorm:"column:dateofbirth;not_null"`
	EmpNo           string    `gorm:"column:empno;not_null"`
	IsActive        bool      `gorm:"column:isactive;not_null;default:false"`
	IsDisables      bool      `gorm:"column:isdisables;not_null;default:false"`
	IsVirtual       bool      `gorm:"column:isvirtual;not_null;default:false"`
	IsVirtualActive bool      `gorm:"column:isvirtualactive;not_null;default:false"`
	IsConfirmed     bool      `gorm:"column:isconfirmed;not_null;default:false"`
	Hours           float32   `gorm:"column:hours;not_null;default:0.00"`
	IsShiftWorker   bool      `gorm:"column:isshiftworker;not_null;default:false"`
	IsResigned      bool      `gorm:"column:isresigned;not_null;default:false"`
	AppointmentDate time.Time `gorm:"column:appointmentdate"`
	ResignDate      time.Time `gorm:"column:resigndate"`
}

func (c TableEmployee) TableName() string {
	return "table_employee"
}
