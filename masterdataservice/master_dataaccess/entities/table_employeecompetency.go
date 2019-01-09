package entities

import "github.com/jinzhu/gorm"

type TableEmployeeCompetency struct {
	gorm.Model
	EmployeeId       uint `gorm:"column:employeeid;not_null"`
	CompetencyTypeId uint `gorm:"column:competencytypeid;not_null"`
	Employee         TableEmployee
	CompetencyType   TableCompetencyType
}

func (c TableEmployeeCompetency) TableName() string {
	return "table_employeecompetency"
}
