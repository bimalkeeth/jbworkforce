package entities

import "github.com/jinzhu/gorm"

type TableRosterType struct {
	gorm.Model
	RosterTypeName       string  `gorm:"column:rostertypename;type:varchar(200);not_null"`
	RosterTypeAbbr       string  `gorm:"column:rostertypeabbr;type:varchar(50);not_null"`
	HasAnnualLeaveCredit bool    `gorm:"column:hasannualleavecredit;not_null"`
	Colour               string  `gorm:"column:colour;type:varchar(50);not_null"`
	IsActive             bool    `gorm:"column:isactive;not_null"`
	IsAvailableToWork    bool    `gorm:"column:isavailabletowork;not_null"`
	DisplayControl       string  `gorm:"column:displaycontrol;type:varchar(150)"`
	DisplayOrder         float64 `gorm:"column:displayorder;default:0.0"`
}

func (c TableRosterType) TableName() string {
	return "table_rostertype"
}
