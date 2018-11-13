package entities

import "github.com/jinzhu/gorm"

type TableRosterShift struct {
	gorm.Model
	RosterUnitTemplateId uint
	RosterUnitId         uint
	ShiftTypeId          uint
	ClinicalStreamId     uint
	HealthCategoryId     uint
	IsMedical            bool
	IsGenerated          bool
	ShiftLength          float64
	ShiftHour            float64
	WeekClose            float64
	BedLength            float64
	IsADOBackFilled      bool
	IsALBackFilled       bool
	IsInCharge           bool
	SundayEft            float64
	SundayBed            float64
	MondayEft            float64
	MondayBed            float64
	TuesdayEft           float64
	TuesdayBed           float64
	WednesdayEft         float64
	WednesdayBed         float64
	ThursdayEft          float64
	ThursdayBed          float64
	FridayEft            float64
	FridayBed            float64
	SaturdayEft          float64
	SaturdayBed          float64
	WeeklyEft            float64
	WeeklyBed            float64
	Comment              string
	StaffHour            float64
	BedHour              float64
	PMShiftStaff         float64
	PMShiftEft           float64
	NightShiftStaff      float64
	NightShiftEft        float64
	SaturdayShiftEft     float64
	SundayShiftEft       float64
}
