package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type TableEmployeeTransition struct {
	gorm.Model
	EmployeeId        uint           `gorm:"column:employeeid;not_null"`
	PositionName      string         `gorm:"column:positionname"`
	TransitionStatus  string         `gorm:"column:transitionstatus"`
	TransitionComment string         `gorm:"column:transitioncomment"`
	StaffMigrationTag postgres.Jsonb `gorm:"column:staffmigrationtag"`
	TableEmployee     TableEmployee  `gorm:"foreignkey:employeeid"`
}

func (c TableEmployeeTransition) TableName() string {
	return "table_employeetransition"
}
