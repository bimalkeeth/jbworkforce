package entities

import (
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
)

type TableEmployeeTransition struct {
	gorm.Model
	EmployeeId        uint           `gorm:"column:employeeid;not_null"`
	PositionName      string         `gorm:"column:positionname;type:varchar(200)"`
	TransitionStatus  string         `gorm:"column:transitionstatus;type:varchar(100)"`
	TransitionComment string         `gorm:"column:transitioncomment;type:varchar(250)"`
	StaffMigrationTag postgres.Jsonb `gorm:"column:staffmigrationtag"`
	Employee          TableEmployee  `gorm:"foreignkey:employeeid"`
}

func (c TableEmployeeTransition) TableName() string {
	return "table_employeetransition"
}
