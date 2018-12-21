package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

//--------------------------------------------
// Table for Address Type
//--------------------------------------------
type TableAddressType struct {
	gorm.Model
	AddressType string `gorm:"column:addresstype;type:varchar(25);not_null"`
	Description string `gorm:"column:description;type:varchar(300)"`
	Addresses   []TableAddress
}

func (c TableAddressType) TableName() string {
	return "table_addresstype"
}

func (c TableAddressType) Validate(db *gorm.DB) {
	if len(c.AddressType) > 25 {
		_ = db.AddError(errors.New("address type should be less or equel to 25"))
	}
	if len(c.Description) > 200 {
		_ = db.AddError(errors.New("description should be less or equel to 25"))
	}
}
