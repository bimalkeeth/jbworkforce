package entities

import "github.com/jinzhu/gorm"

type TableAgencyContacts struct {
	gorm.Model
	AgencyId      uint          `gorm:"column:agencyid;not_null"`
	ContactId     uint          `gorm:"column:contactid;not_null"`
	TableAgency   TableAgency   `gorm:"foreignkey:agencyid"`
	TableContacts TableContacts `gorm:"foreignkey:contactid"`
}

func (c TableAgencyContacts) TableName() string {
	return "table_agencycontacts"
}
