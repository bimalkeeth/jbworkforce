package entities

import "github.com/jinzhu/gorm"

type TableCompetencyType struct {
	gorm.Model
	CompentencyName    string `gorm:"column:compentencyname;type:varchar(200);not_null"`
	WeekLength         int    `gorm:"column:weeklength;not_null"`
	CompentencyAbbr    string `gorm:"column:compentencyabbr;not_null"`
	EmployeeCompetency []TableEmployeeCompetency
}

func (c TableCompetencyType) TableName() string {
	return "table_compentencytype"
}
