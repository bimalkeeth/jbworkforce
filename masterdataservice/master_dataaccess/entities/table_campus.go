package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableCampus struct {
	gorm.Model
	CampusName              string    `gorm:"column:campusname;type:varchar(100);not_null"`
	CampusAbbr              string    `gorm:"column:campusabbr;type:varchar(50);not_null"`
	Description             string    `gorm:"column:description;type:varchar(400)"`
	YearlyPlannerDate       time.Time `gorm:"column:yearlyplannerdate"`
	YearlyPlannerPeriodDate time.Time `gorm:"column:yearlyplannerperioddate"`
	IsActive                bool      `gorm:"column:isactive;not_null;default:true"`
	Divisions               []TableDivision
}

func (c TableCampus) TableName() string {
	return "table_campus"
}
