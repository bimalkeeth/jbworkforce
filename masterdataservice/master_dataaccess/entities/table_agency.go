package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableAgency struct {
	gorm.Model
	AgencyName      string `gorm:"column:agencyname;type:varchar(200);not_null"`
	AgencyAddresses []*TableAgencyAddress
	AgencyContacts  []*TableAgencyContacts
}

func (c TableAgency) TableName() string {
	return "table_agency"
}

func (c TableAgency) Validate(db *gorm.DB) {
	if len(c.AgencyName) > 200 {
		_ = db.AddError(errors.New("agency name should be less or equal to 25"))
	}
	if len(c.AgencyName) == 0 {
		_ = db.AddError(errors.New("agency name should not be empty"))
	}
}
