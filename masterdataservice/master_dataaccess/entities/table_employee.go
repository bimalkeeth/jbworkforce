package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type TableEmployee struct {
	gorm.Model
	FirstName       string    `gorm:"column:firstname;type:varchar(200);not_null"`
	LastName        string    `gorm:"column:lastname;type:varchar(200);not_null"`
	PrefName        string    `gorm:"column:name;type:varchar(150);not_null"`
	GenderId        uint      `gorm:"column:genderid;not_null"`
	DateOfBirth     time.Time `gorm:"column:dateofbirth;not_null"`
	EmpNo           string    `gorm:"column:empno;type:varchar(150);not_null"`
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

	EmployeeDetail               *TableEmployeeDetail
	EmployeeContacts             []*TableEmployeeContact
	EmployeeAddress              []*TableEmployeeAddress
	EmployeeAgencies             []*TableEmployeeAgency
	EmployeeClinicalStream       *TableEmployeeClinicalStream
	EmployeeAllowance            []*TableEmployeeAllowance
	EmployeeCompetency           *TableEmployeeCompetency
	EmployeeTransitions          []*TableEmployeeTransition
	EmployeeDepartmentPreference []*TableEmployeeDepartmentPreference
	EmployeeDiscrepancies        []*TableEmployeeDiscrepancy
	EmployeeEmpType              *TableEmployeeEmpType
	EmployeeGroupEmp             *TableEmployeeGroupEmp
	Gender                       *TableGender
}

func (c TableEmployee) TableName() string {
	return "table_employee"
}

func (c TableEmployee) Validate(db *gorm.DB) {
	if len(c.FirstName) > 200 {
		_ = db.AddError(errors.New("name length should be less or equal to 200"))
	}
	if len(c.LastName) > 200 {
		_ = db.AddError(errors.New("last name length should be less or equal to 200"))
	}
	if len(c.PrefName) > 200 {
		_ = db.AddError(errors.New("pref. name length should be less or equal to 150"))
	}
	if len(c.EmpNo) > 150 {
		_ = db.AddError(errors.New("empno length should be less or equal to 150"))
	}
}
