package entities

import "github.com/jinzhu/gorm"

type TableState struct {
	gorm.Model
	StateName string `gorm:"column:statename;type:varchar(200);not_null"`
	StateAbbr string `gorm:"column:stateabbr;type:varchar(50);not_null"`
}

func (c TableState) TableName() string {
	return "table_state"
}
