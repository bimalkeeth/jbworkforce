package entities

import "github.com/jinzhu/gorm"

//--------------------------------------------
// Table for Address
//--------------------------------------------
type TableAddress struct {
	gorm.Model
	Street           string `gorm:"column:street;not_null"`
	Street2          string
	City             string           `gorm:"column:city;not_null"`
	PostCode         string           `gorm:"column:postcode;not_null"`
	StateId          uint             `gorm:"column:stateid;not_null"`
	Country          string           `gorm:"column:country"`
	AddressTypeId    uint             `gorm:"column:addresstypeid;not_null"`
	TableAddressType TableAddressType `gorm:"foreignkey:addresstypeid"`
}

func (c TableAddress) TableName() string {
	return "table_address"
}
