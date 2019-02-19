package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

//--------------------------------------------
// Table for Address
//--------------------------------------------
type TableAddress struct {
	gorm.Model
	Street        string `gorm:"column:street;type:varchar(200);not_null"`
	Street2       string `gorm:"column:street2;type:varchar(200)"`
	City          string `gorm:"column:city;type:varchar(100);not_null"`
	PostCode      string `gorm:"column:postcode;type:varchar(25);not_null"`
	StateId       uint   `gorm:"column:stateid;not_null"`
	Country       string `gorm:"column:country;type:varchar(100)"`
	AddressTypeId uint   `gorm:"column:addresstypeid;not_null"`
	AddressType   *TableAddressType
	State         *TableState
}

func (c TableAddress) TableName() string {
	return "table_address"
}

func (c TableAddress) Validate(db *gorm.DB) {
	if len(c.Street) > 200 {
		_ = db.AddError(errors.New("street should be less or equal to 200"))
	}
	if c.Street == "" {
		_ = db.AddError(errors.New("street field can not be empty"))
	}
	if len(c.Street2) > 200 {
		_ = db.AddError(errors.New("street2 should be less or equal to 200"))
	}
	if len(c.City) > 100 {
		_ = db.AddError(errors.New("city should be less or equal to 100"))
	}
	if c.City == "" {
		_ = db.AddError(errors.New("city should be not empty"))
	}
}
