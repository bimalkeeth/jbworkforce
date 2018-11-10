package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableFinancialYear struct {
	gorm.Model
	FinancialYearName string    `gorm:"column:financialyearname;not_null"`
	StartYear         int       `gorm:"column:startyear;not_null"`
	StartDate         time.Time `gorm:"column:startdate;not_null"`
	EndDate           time.Time `gorm:"column:enddate;not_null"`
}

func (c TableFinancialYear) TableName() string {
	return "table_financialyear"
}
