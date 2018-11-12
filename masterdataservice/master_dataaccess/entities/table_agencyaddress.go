package entities

import "github.com/jinzhu/gorm"

type TableAgencyAddress struct {
	gorm.Model
	AgencyId     uint         `gorm:"column:agencyid;not_null"`
	AddressId    uint         `gorm:"column:addressid;not_null"`
	TableAgency  TableAgency  `gorm:"foreignkey:agencyid"`
	TableAddress TableAddress `gorm:"foreignkey:addressid"`
}

func (c TableAgencyAddress) TableName() string {
	return "table_agencyaddress"
}
