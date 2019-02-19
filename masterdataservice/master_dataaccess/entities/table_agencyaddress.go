package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableAgencyAddress struct {
	gorm.Model
	AgencyId  uint `gorm:"column:agencyid;not_null"`
	AddressId uint `gorm:"column:addressid;not_null"`
	Agency    *TableAgency
	Address   *TableAddress
}

func (c TableAgencyAddress) TableName() string {
	return "table_agencyaddress"
}

func (c TableAgencyAddress) Validate(db *gorm.DB) {
	if c.AgencyId == 0 {
		_ = db.AddError(errors.New("agency id should not be empty"))
	}
	if c.AddressId == 0 {
		_ = db.AddError(errors.New("addressid should not be empty"))
	}
}
