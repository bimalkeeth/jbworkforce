package entities

import "github.com/jinzhu/gorm"

type TablePostcode struct {
	gorm.Model
	Location string `gorm:"column:location;type:varchar(200)"`
	Postcode string `gorm:"column:postcode;type:varchar(50)"`
}

func (c TablePostcode) TableName() string {
	return "table_postcode"
}
