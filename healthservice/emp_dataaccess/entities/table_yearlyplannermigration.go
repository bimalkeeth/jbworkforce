package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableYearlyPlannerMigration struct {
	gorm.Model
	DumpType string    `gorm:"column:dumptype;not_null"`
	DateFrom time.Time `gorm:"column:datefrom;not_null"`
	DateTo   time.Time `gorm:"column:dateto;not_null"`
}

func (c TableYearlyPlannerMigration) TableName() string {
	return "table_yearlyplannermigration"
}
