package entities

import "github.com/jinzhu/gorm"

type TableAgency struct {
	gorm.Model
	AgencyName string `gorm:"column:agencyname;not_null"`
}

func (c TableAgency) TableName() string {
	return "table_agency"
}
