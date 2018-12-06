package entities

import "github.com/jinzhu/gorm"

//--------------------------------------------
// Table for Address Type
//--------------------------------------------
type TableAddressType struct {
	gorm.Model
	AddressType string         `gorm:"column:addresstype;not_null"`
	Description string         `gorm:"column:description"`
	Addresses   []TableAddress `gorm:"foreignkey:addresstypeid"`
}

func (c TableAddressType) TableName() string {
	return "table_addresstype"
}
