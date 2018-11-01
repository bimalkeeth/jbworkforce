package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableCampus struct {
	gorm.Model
	CampusName              string    `gorm:"column:campusname;not_null"`
	CampusAbbr              string    `gorm:"column:campusabbr;not_null"`
	Description             string    `gorm:"column:description"`
	YearlyPlannerDate       time.Time `gorm:"column:yearlyplannerdate"`
	YearlyPlannerPeriodDate time.Time `gorm:"column:yearlyplannerperioddate"`
	IsActive                bool      `gorm:"column:isactive;not_null;default:true"`
}

func (c TableCampus) TableName() string {
	return "table_campus"
}
