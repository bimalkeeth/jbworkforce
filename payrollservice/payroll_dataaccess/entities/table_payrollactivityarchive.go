package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TablePayrollActivityArchive struct {
	gorm.Model
	PayPeriodId         string    `gorm:"column:urnumber"`
	PayPeriodSubId      string    `gorm:"column:urnumber"`
	PayPeriodDate       time.Time `gorm:"column:urnumber"`
	DateWeekId          uint      `gorm:"column:urnumber"`
	EmployeeId          uint      `gorm:"column:urnumber"`
	HealthCategoryId    uint      `gorm:"column:urnumber"`
	SkillTypeId         uint      `gorm:"column:urnumber"`
	ClinicalStreamId    uint      `gorm:"column:urnumber"`
	RosterTypeId        uint      `gorm:"column:urnumber"`
	SubDepartmentId     uint      `gorm:"column:urnumber"`
	DepartmentId        uint      `gorm:"column:urnumber"`
	CostCentreId        uint      `gorm:"column:urnumber"`
	GradeCodeId         uint      `gorm:"column:urnumber"`
	StaffTypeId         uint      `gorm:"column:urnumber"`
	BasicHour           float64   `gorm:"column:urnumber"`
	ExtraHour           float64   `gorm:"column:urnumber"`
	IsImport            bool      `gorm:"column:urnumber"`
	EmployeeNumber      string    `gorm:"column:urnumber"`
	EmployeeFirstName   string    `gorm:"column:urnumber"`
	EmployeeLastName    string    `gorm:"column:urnumber"`
	AppointmentType     string    `gorm:"column:urnumber"`
	ContractedHour      float64   `gorm:"column:urnumber"`
	EFT                 float64   `gorm:"column:urnumber"`
	GradeCode           string    `gorm:"column:urnumber"`
	GradeCodeAbbr       string    `gorm:"column:urnumber"`
	GradeSubCode        string    `gorm:"column:urnumber"`
	GradeDescription    string    `gorm:"column:urnumber"`
	PositionId          string    `gorm:"column:urnumber"`
	PositionDescription string    `gorm:"column:urnumber"`
	UnitId              string    `gorm:"column:urnumber"`
	UnitName            string    `gorm:"column:urnumber"`
	PayTypeGroupId      string    `gorm:"column:urnumber"`
	PayTypeGroup        string    `gorm:"column:urnumber"`
	PayType             string    `gorm:"column:urnumber"`
	PayTypeDescription  string    `gorm:"column:urnumber"`
	PayTypeHour         float64   `gorm:"column:urnumber"`
	AmountPaid          float64   `gorm:"column:urnumber"`
	Note                string    `gorm:"column:urnumber"`
}

func (c TablePayrollActivityArchive) TableName() string {
	return "table_payrollactivityarchive"
}
