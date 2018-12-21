package entities

import "github.com/jinzhu/gorm"

type TableAgency struct {
	gorm.Model
	AgencyName      string `gorm:"column:agencyname;type:varchar(200);not_null"`
	AgencyAddresses []TableAgencyAddress
	AgencyContacts  []TableAgencyContacts
}

func (c TableAgency) TableName() string {
	return "table_agency"
}
