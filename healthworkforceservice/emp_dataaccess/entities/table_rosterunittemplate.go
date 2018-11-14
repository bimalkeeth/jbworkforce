package entities

import "github.com/jinzhu/gorm"

type TableRosterUnitTemplate struct {
	gorm.Model
	RosterUnitId          uint
	ShiftTypeId           uint
	ClinicalStreamId      uint
	HealthCategoryId      uint
	TotalBeds             int
	Occupancy             float64
	DaysOpen              int
	TotalStaffRequired    int
	TotalPersonInCharge   int
	IsInChargePartOfShift bool
	StaffToBedRatio       float64
	ShiftLength           float64
	IsADOBackFilled       bool
	IsALBackFilled        bool
	IsSpecialist          bool
}
