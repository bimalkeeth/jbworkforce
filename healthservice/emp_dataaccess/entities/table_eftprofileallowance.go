package entities

import "github.com/jinzhu/gorm"

type TableEftProfileAllowance struct {
	gorm.Model
	EftProfileId    uint            `gorm:"column:eftprofileid;not_null"`
	AllowanceId     uint            `gorm:"column:allowanceid;not_null"`
	Name            string          `gorm:"column:name;not_null"`
	Description     string          `gorm:"column:eftprofileactivityId;not_null"`
	Amount          float64         `gorm:"column:amount;not_null;default:0.0"`
	TableEftProfile TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableEftProfileAllowance) TableName() string {
	return "table_eftprofileallowance"
}
