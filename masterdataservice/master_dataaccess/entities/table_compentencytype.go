package entities

import "github.com/jinzhu/gorm"

type TableCompetencyType struct {
	gorm.Model
	CompentencyName    string                    `gorm:"column:compentencyname;not_null"`
	WeekLength         int                       `gorm:"column:weeklength;not_null"`
	CompentencyAbbr    string                    `gorm:"column:compentencyabbr;not_null"`
	EmployeeCompetency []TableEmployeeCompetency `gorm:"foreignkey:competencytypeid"`
}

func (c TableCompetencyType) TableName() string {
	return "table_compentencytype"
}
