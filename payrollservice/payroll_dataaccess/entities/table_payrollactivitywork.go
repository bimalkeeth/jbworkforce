package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TablePayrollActivityWork struct {
	gorm.Model
	PayPeriodId         string    `gorm:"column:payperiodid"`
	PayPeriodSubId      string    `gorm:"column:payperiodsubid"`
	PayPeriodDate       time.Time `gorm:"column:payperioddate"`
	DateWeekId          uint      `gorm:"column:dateweekid"`
	EmployeeId          uint      `gorm:"column:employeeid"`
	HealthCategoryId    uint      `gorm:"column:healthcategoryid"`
	SkillTypeId         uint      `gorm:"column:skilltypeid"`
	ClinicalStreamId    uint      `gorm:"column:clinicalstreamid"`
	RosterTypeId        uint      `gorm:"column:rostertypeid"`
	SubDepartmentId     uint      `gorm:"column:subdepartmentid"`
	DepartmentId        uint      `gorm:"column:departmentid"`
	CostCentreId        uint      `gorm:"column:costcentreid"`
	GradeCodeId         uint      `gorm:"column:gradecodeid"`
	StaffTypeId         uint      `gorm:"column:stafftypeid"`
	BasicHour           float64   `gorm:"column:basichour"`
	ExtraHour           float64   `gorm:"column:extrahour"`
	IsImport            bool      `gorm:"column:isimport"`
	EmployeeNumber      string    `gorm:"column:employeenumber"`
	EmployeeFirstName   string    `gorm:"column:employeefirstname"`
	EmployeeLastName    string    `gorm:"column:employeelastname"`
	AppointmentType     string    `gorm:"column:appointmenttype"`
	ContractedHour      float64   `gorm:"column:contractedhour"`
	EFT                 float64   `gorm:"column:eft"`
	GradeCode           string    `gorm:"column:gradecode"`
	GradeCodeAbbr       string    `gorm:"column:gradecodeabbr"`
	GradeSubCode        string    `gorm:"column:gradesubcode"`
	GradeDescription    string    `gorm:"column:gradedescription"`
	PositionId          string    `gorm:"column:positionid"`
	PositionDescription string    `gorm:"column:positiondescription"`
	UnitId              string    `gorm:"column:unitid"`
	UnitName            string    `gorm:"column:unitname"`
	PayTypeGroupId      string    `gorm:"column:paytypegroupid"`
	PayTypeGroup        string    `gorm:"column:paytypegroup"`
	PayType             string    `gorm:"column:paytype"`
	PayTypeDescription  string    `gorm:"column:paytypedescription"`
	PayTypeHour         float64   `gorm:"column:paytypehour"`
	AmountPaid          float64   `gorm:"column:amountpaid"`
	Note                string    `gorm:"column:note"`
}

func (c TablePayrollActivityWork) TableName() string {
	return "table_payrollactivitywork"
}
