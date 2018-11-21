package entities

import "github.com/jinzhu/gorm"

type TableEftProfileSpread struct {
	gorm.Model
	EftProfileId      uint            `gorm:"column:eftprofileid;not_null"`
	DateWeekId        uint            `gorm:"column:dateweekid;not_null"`
	DayNumber         int             `gorm:"column:daynumber;not_null"`
	DailySpread       float64         `gorm:"column:dailyspread;default:0.0"`
	HolidaySpread     float64         `gorm:"column:holidayspread;default:0.0"`
	WeekendSpread     float64         `gorm:"column:weekendspread;default:0.0"`
	ClosurePercentage float64         `gorm:"column:closurepercentage;default:0.0"`
	IsInput           bool            `gorm:"column:isinput;default:false"`
	IsActive          bool            `gorm:"column:isactive;default:true"`
	Description       string          `gorm:"column:description"`
	TableEftProfile   TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableEftProfileSpread) TableName() string {
	return "table_eftprofilespread"
}
