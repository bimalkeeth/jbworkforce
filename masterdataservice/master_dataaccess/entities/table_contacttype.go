package entities

import "github.com/jinzhu/gorm"

type TableContactType struct {
	gorm.Model
	ContactType string `gorm:"column:contacttype;type:varchar(100);not_null"`
	Description string `gorm:"column:description;type:varchar(400)"`
	Contacts    []TableContacts
}

func (c TableContactType) TableName() string {
	return "table_contacttype"
}
