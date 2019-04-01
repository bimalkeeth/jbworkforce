package entities

import (
	"errors"
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
	Divisions               []*TableDivision
}

func (c TableCampus) TableName() string {
	return "table_campus"
}
func (c TableCampus) Validate(db *gorm.DB) {

	if len(c.CampusName) > 200 {
		_ = db.AddError(errors.New("campus name length should be less or equal to 128"))
	}
	if len(c.CampusAbbr) > 50 {
		_ = db.AddError(errors.New("campus abbr length should be less or equal to 50"))
	}
	if len(c.Description) > 400 {
		_ = db.AddError(errors.New("description length should be less or equal to 400"))
	}
}
