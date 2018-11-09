package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableDateWeek struct {
	gorm.Model
	FinancialYearId  uint      `gorm:"column:financialyearid;not_null"`
	EachDay          time.Time `gorm:"column:eachday;not_null"`
	EachDayName      string    `gorm:"column:eachdayname;not_null"`
	YearlyWeekNumber int       `gorm:"column:yearlyweeknumber;not_null"`
	IsHoliday        bool      `gorm:"column:isholiday;not_null"`
	IsHolidayBudget  bool      `gorm:"column:isholidaybudget;not_null"`
	WeekPeriod       int       `gorm:"column:weekperiod;not_null"`
	PayPeriod        int       `gorm:"column:payperiod;not_null"`
	DateName         string    `gorm:"column:datename;not_null"`
	Week             int       `gorm:"column:week;not_null"`
	WeekName         string    `gorm:"column:weekname;not_null"`
	Month            int       `gorm:"column:month;not_null"`
	MonthName        string    `gorm:"column:monthname;not_null"`
	MonthNumber      string    `gorm:"column:monthnumber;not_null"`
	Year             int       `gorm:"column:year;not_null"`
	DaysInMonth      int       `gorm:"column:daysinmonth;not_null"`
}

func (c TableDateWeek) TableName() string {
	return "table_dateweek"
}
