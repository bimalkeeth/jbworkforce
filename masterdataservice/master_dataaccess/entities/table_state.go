package entities

import "github.com/jinzhu/gorm"

type TableState struct {
	gorm.Model
	StateName string `gorm:"column:statename;not_null"`
	StateAbbr string `gorm:"column:stateabbr;not_null"`
}

func (c TableState) TableName() string {
	return "table_state"
}
