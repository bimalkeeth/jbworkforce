package entities

import "github.com/jinzhu/gorm"

type TableAgencyAddress struct {
	gorm.Model
	AgencyId  uint `gorm:"column:agencyid;not_null"`
	AddressId uint `gorm:"column:addressid;not_null"`
	Agency    TableAgency
	Address   TableAddress
}

func (c TableAgencyAddress) TableName() string {
	return "table_agencyaddress"
}
