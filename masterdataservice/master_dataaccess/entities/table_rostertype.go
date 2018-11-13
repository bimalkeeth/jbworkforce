package entities

import "github.com/jinzhu/gorm"

type TableRosterType struct {
	gorm.Model
	RosterTypeName       string  `gorm:"column:rostertypename;not_null"`
	RosterTypeAbbr       string  `gorm:"column:rostertypeabbr;not_null"`
	HasAnnualLeaveCredit bool    `gorm:"column:hasannualleavecredit;not_null"`
	Colour               string  `gorm:"column:colour;not_null"`
	IsActive             bool    `gorm:"column:isactive;not_null"`
	IsAvailableToWork    bool    `gorm:"column:isavailabletowork;not_null"`
	DisplayControl       string  `gorm:"column:displaycontrol"`
	DisplayOrder         float64 `gorm:"column:displayorder;default:0.0"`
}

func (c TableRosterType) TableName() string {
	return "table_rostertype"
}
