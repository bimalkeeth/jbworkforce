package entities

import "github.com/jinzhu/gorm"

type TableEftProfileActivityMapping struct {
	gorm.Model
	EftProfileActivityId    uint                    `gorm:"column:eftprofileactivityId;not_null"`
	RosterUnitId            uint                    `gorm:"column:rosterunitId;not_null"`
	TableEftProfileActivity TableEftProfileActivity `gorm:"foreignkey:eftprofileactivityid"`
	TableRosterUnit         TableRosterUnit         `gorm:"foreignkey:rosterunitId"`
}

func (c TableEftProfileActivityMapping) TableName() string {
	return "table_eftprofileactivitymapping"
}
