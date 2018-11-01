package entities

import "github.com/jinzhu/gorm"

type TableDivision struct {
	gorm.Model
	DivisionName string      `gorm:"column:divisionname;not_null"`
	DivisionAbbr string      `gorm:"column:divisionabbr;not_null"`
	Description  string      `gorm:"column:description"`
	CampusId     uint        `gorm:"column:campusid;not_null"`
	TableCampus  TableCampus `gorm:"foreignkey:campusid"`
}

func (c TableDivision) TableName() string {
	return "table_division"
}
