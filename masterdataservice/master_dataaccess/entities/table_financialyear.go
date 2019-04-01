package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type TableFinancialYear struct {
	gorm.Model
	FinancialYearName string    `gorm:"column:financialyearname;type:varchar(200);not_null"`
	StartYear         int       `gorm:"column:startyear;not_null"`
	StartDate         time.Time `gorm:"column:startdate;not_null"`
	EndDate           time.Time `gorm:"column:enddate;not_null"`
}

func (c TableFinancialYear) TableName() string {
	return "table_financialyear"
}

func (c TableFinancialYear) Validate(db *gorm.DB) {
	if len(c.FinancialYearName) > 200 {
		_ = db.AddError(errors.New("financial year name should be less or equal to 100"))
	}
}
