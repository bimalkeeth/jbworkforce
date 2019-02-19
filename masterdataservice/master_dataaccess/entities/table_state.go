package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableState struct {
	gorm.Model
	StateName string `gorm:"column:statename;type:varchar(200);not_null"`
	StateAbbr string `gorm:"column:stateabbr;type:varchar(50);not_null"`
}

func (c TableState) TableName() string {
	return "table_state"
}

//--------------------------------------------------
// Validate Fields of State Table
//--------------------------------------------------
func (c TableState) Validate(db *gorm.DB) {
	if len(c.StateName) > 200 {
		_ = db.AddError(errors.New("statename should be less or equal to 200"))
	}
	if c.StateName == "" {
		_ = db.AddError(errors.New("statename field can not be empty"))
	}
	if len(c.StateAbbr) > 200 {
		_ = db.AddError(errors.New("stateabbr should be less or equal to 50"))
	}
	if c.StateAbbr == "" {
		_ = db.AddError(errors.New("stateabbr should be not empty"))
	}
}
