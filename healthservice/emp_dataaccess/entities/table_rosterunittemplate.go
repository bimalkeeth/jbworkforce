package entities

import "github.com/jinzhu/gorm"

type TableRosterUnitTemplate struct {
	gorm.Model
	RosterUnitId          uint    `gorm:"column:rosterunitid;not_null"`
	ShiftTypeId           uint    `gorm:"column:shifttypeid;not_null"`
	ClinicalStreamId      uint    `gorm:"column:clinicalstreamid;not_null"`
	HealthCategoryId      uint    `gorm:"column:healthcategoryid;not_null"`
	TotalBeds             int     `gorm:"column:totalbeds;default:0"`
	Occupancy             float64 `gorm:"column:occupancy;default:0.0"`
	DaysOpen              int     `gorm:"column:daysopen;default:0"`
	TotalStaffRequired    int     `gorm:"column:totalstaffrequired;default:0"`
	TotalPersonInCharge   int     `gorm:"column:totalpersonincharge;default:0"`
	IsInChargePartOfShift bool    `gorm:"column:isinchargepartofshift;default:false"`
	StaffToBedRatio       float64 `gorm:"column:stafftobedratio;default:0.0"`
	ShiftLength           float64 `gorm:"column:shiftlength;default:0.0"`
	IsADOBackFilled       bool    `gorm:"column:isadobackfilled;default:false"`
	IsALBackFilled        bool    `gorm:"column:isalbackfilled;default:false"`
	IsSpecialist          bool    `gorm:"column:isspecialist;default:false"`

	TableRosterUnit TableRosterUnit `gorm:"foreignkey:rosterunitid"`
	TableCatStream  TableCatStream  `gorm:"foreignkey:clinicalstreamid"`
}

func (c TableRosterUnitTemplate) TableName() string {
	return "table_rosterunittemplate"
}
