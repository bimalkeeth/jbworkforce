package entities

import "github.com/jinzhu/gorm"

type TableEmployeeCompetency struct {
	gorm.Model
	EmployeeId          uint                `gorm:"column:employeeid;not_null"`
	CompetencyTypeId    uint                `gorm:"column:competencytypeid;not_null"`
	TableEmployee       TableEmployee       `gorm:"foreignkey:employeeid"`
	TableCompetencyType TableCompetencyType `gorm:"foreignkey:competencytypeid"`
}

func (c TableEmployeeCompetency) TableName() string {
	return "table_employeecompetency"
}
