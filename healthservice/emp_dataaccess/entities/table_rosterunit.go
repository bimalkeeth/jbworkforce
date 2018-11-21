package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type TableRosterUnit struct {
	gorm.Model
	EftProfileId          uint            `gorm:"column:eftprofileid;not_null"`
	HealthCategoryId      uint            `gorm:"column:healthcategoryid;not_null"`
	ClinicalStreamId      uint            `gorm:"column:clinicalstreamid;not_null"`
	SecondaryCostCentreId uint            `gorm:"column:secondarycostcentreid;not_null"`
	RosterUnitType        string          `gorm:"column:rosterunittype;not_null"`
	RosterUnitOrder       int             `gorm:"column:rosterunitorder;default:0"`
	RosterUnitName        string          `gorm:"column:rosterunitname;not_null"`
	Description           string          `gorm:"column:description"`
	BedOpenHour           float64         `gorm:"column:BedOpenHour;default:0.0"`
	ShiftLengthAdjustment string          `gorm:"column:shiftlengthadjustment"`
	Setting               postgres.Jsonb  `gorm:"column:setting"`
	TableEftProfile       TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableRosterUnit) TableName() string {
	return "table_rosterunit"
}
