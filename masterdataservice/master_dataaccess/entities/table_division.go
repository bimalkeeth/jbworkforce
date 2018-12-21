package entities

import "github.com/jinzhu/gorm"

type TableDivision struct {
	gorm.Model
	DivisionName string `gorm:"column:divisionname;type:varchar(200);not_null"`
	DivisionAbbr string `gorm:"column:divisionabbr;type:varchar(50);not_null"`
	Description  string `gorm:"column:description;type:varchar(400)"`
	CampusId     uint   `gorm:"column:campusid;not_null"`
	Campus       TableCampus
	Departments  []TableDepartment
}

func (c TableDivision) TableName() string {
	return "table_division"
}
