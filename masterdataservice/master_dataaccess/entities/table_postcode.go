package entities

import "github.com/jinzhu/gorm"

type TablePostcode struct {
	gorm.Model
	Location string `gorm:"column:location"`
	Postcode string `gorm:"column:postcode"`
}

func (c TablePostcode) TableName() string {
	return "table_postcode"
}
