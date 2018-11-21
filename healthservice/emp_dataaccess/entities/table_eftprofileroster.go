package entities

import "github.com/jinzhu/gorm"

type TableEftProfileRoster struct {
	gorm.Model
	EftProfileId      uint    `gorm:"column:eftprofileid;not_null"`
	DateWeekId        uint    `gorm:"column:dateweekid"`
	DayNumber         int     `gorm:"column:eftprofileid;not_null"`
	DailySpread       float64 `gorm:"column:eftprofileid;not_null"`
	HolidaySpread     float64 `gorm:"column:eftprofileid;not_null"`
	WeekendSpread     float64 `gorm:"column:eftprofileid;not_null"`
	ClosurePercentage float64 `gorm:"column:eftprofileid;not_null"`
	IsInput           bool    `gorm:"column:eftprofileid;not_null"`
	IsActive          bool    `gorm:"column:eftprofileid;not_null"`
	Description       string  `gorm:"column:eftprofileid;not_null"`
}

func (c TableEftProfileRoster) TableName() string {
	return "table_eftprofileroster"
}
