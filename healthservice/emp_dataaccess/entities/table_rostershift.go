package entities

import "github.com/jinzhu/gorm"

type TableRosterShift struct {
	gorm.Model
	RosterUnitTemplateId uint    `gorm:"column:rosterunittemplateid"`
	RosterUnitId         uint    `gorm:"column:rosterunitid;not_null"`
	ShiftTypeId          uint    `gorm:"column:shifttypeid;not_null"`
	ClinicalStreamId     uint    `gorm:"column:clinicalstreamid"`
	HealthCategoryId     uint    `gorm:"column:healthcategoryid"`
	IsMedical            bool    `gorm:"column:ismedical;default:false"`
	IsGenerated          bool    `gorm:"column:isgenerated;default:false"`
	ShiftLength          float64 `gorm:"column:shiftlength;default:0.0"`
	ShiftHour            float64 `gorm:"column:shifthour;default:0.0"`
	WeekClose            float64 `gorm:"column:weekclose;default:0.0"`
	BedLength            float64 `gorm:"column:bedlength;default:0.0"`
	IsADOBackFilled      bool    `gorm:"column:isadobackfilled;default:false"`
	IsALBackFilled       bool    `gorm:"column:isalbackfilled;default:false"`
	IsInCharge           bool    `gorm:"column:isincharge;default:false"`
	SundayEft            float64 `gorm:"column:sundayeft;default:0.0"`
	SundayBed            float64 `gorm:"column:sundaybed;default:0.0"`
	MondayEft            float64 `gorm:"column:mondayeft;default:0.0"`
	MondayBed            float64 `gorm:"column:mondaybed;default:0.0"`
	TuesdayEft           float64 `gorm:"column:tuesdayeft;default:0.0"`
	TuesdayBed           float64 `gorm:"column:tuesdaybed;default:0.0"`
	WednesdayEft         float64 `gorm:"column:wednesdayeft;default:0.0"`
	WednesdayBed         float64 `gorm:"column:wednesdaybed;default:0.0"`
	ThursdayEft          float64 `gorm:"column:thursdayeft;default:0.0"`
	ThursdayBed          float64 `gorm:"column:thursdaybed;default:0.0"`
	FridayEft            float64 `gorm:"column:fridayeft;default:0.0"`
	FridayBed            float64 `gorm:"column:fridaybed;default:0.0"`
	SaturdayEft          float64 `gorm:"column:saturdayeft;default:0.0"`
	SaturdayBed          float64 `gorm:"column:saturdaybed;default:0.0"`
	WeeklyEft            float64 `gorm:"column:weeklyeft;default:0.0"`
	WeeklyBed            float64 `gorm:"column:weeklybed;default:0.0"`
	Comment              string  `gorm:"column:comment"`
	StaffHour            float64 `gorm:"column:staffhour;default:0.0"`
	BedHour              float64 `gorm:"column:bedhour;default:0.0"`
	PMShiftStaff         float64 `gorm:"column:pmshiftstaff;default:0.0"`
	PMShiftEft           float64 `gorm:"column:pmshifteft;default:0.0"`
	NightShiftStaff      float64 `gorm:"column:nightshiftstaff;default:0.0"`
	NightShiftEft        float64 `gorm:"column:nightshifteft;default:0.0"`
	SaturdayShiftEft     float64 `gorm:"column:saturdayshifteft;default:0.0"`
	SundayShiftEft       float64 `gorm:"column:sundayshifteft;default:0.0"`

	TableRosterUnit         TableRosterUnit         `gorm:"foreignkey:rosterunitid"`
	TableRosterUnitTemplate TableRosterUnitTemplate `gorm:"foreignkey:rosterunittemplateid"`
}

func (c TableRosterShift) TableName() string {
	return "table_rostershift"
}
