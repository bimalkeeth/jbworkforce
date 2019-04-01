package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableDivision struct {
	gorm.Model
	DivisionName string `gorm:"column:divisionname;type:varchar(200);not_null"`
	DivisionAbbr string `gorm:"column:divisionabbr;type:varchar(50);not_null"`
	Description  string `gorm:"column:description;type:varchar(400)"`
	CampusId     uint   `gorm:"column:campusid;not_null"`
	Campus       *TableCampus
	Departments  []*TableDepartment
}

func (c TableDivision) TableName() string {
	return "table_division"
}

func (c TableDivision) Validate(db *gorm.DB) {

	if len(c.DivisionName) > 200 {
		_ = db.AddError(errors.New("division name length should be less or equal to 128"))
	}
	if len(c.DivisionAbbr) > 50 {
		_ = db.AddError(errors.New("division abbr length should be less or equal to 50"))
	}
	if len(c.Description) > 400 {
		_ = db.AddError(errors.New("description length should be less or equal to 400"))
	}
}
