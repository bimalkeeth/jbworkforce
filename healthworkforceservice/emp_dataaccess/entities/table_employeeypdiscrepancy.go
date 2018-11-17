package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableStaffYpDiscrepancy struct {
	gorm.Model
	EmployeeId          uint      `gorm:"column:employeeid;not_null"`
	ReportDate          time.Time `gorm:"column:reportdate"`
	DepartmentId        uint      `gorm:"column:departmentid;not_null"`
	DiscrepancyType     string    `gorm:"column:discrepancytype;not_null"`
	Status              string    `gorm:"column:status;not_null"`
	StartDateWeekId     uint      `gorm:"column:startdateweekid;not_null"`
	EndDateWeekId       uint      `gorm:"column:enddateweekid;not_null"`
	SubDepartmentId     uint      `gorm:"column:subdepartmentid;not_null"`
	ClinicalStreamId    uint      `gorm:"column:clinicalstreamid"`
	Hour                float64   `gorm:"column:hour;not_null"`
	ResignationDate     time.Time `gorm:"column:resignationdate"`
	OldSubDepartmentId  uint      `gorm:"column:oldsubdepartmentid"`
	OldClinicalStreamId uint      `gorm:"column:oldclinicalstreamid"`
	OldHour             float64   `gorm:"column:oldhour"`
	OldResignationDate  time.Time `gorm:"column:oldresignationdate"`
	Description         string    `gorm:"column:description"`
	IsUserChange        bool      `gorm:"column:isuserchange"`
	IsReported          bool      `gorm:"column:IsReported"`
	UserId              uint      `gorm:"column:userid;not_null"`
}

func (c TableStaffYpDiscrepancy) TableName() string {
	return "table_staffypdiscrepancy"
}
