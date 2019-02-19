package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableContactType struct {
	gorm.Model
	ContactType string `gorm:"column:contacttype;type:varchar(25);not_null"`
	Description string `gorm:"column:description;type:varchar(400)"`
	Contacts    []TableContacts
}

func (c TableContactType) TableName() string {
	return "table_contacttype"
}

func (c TableContactType) Validate(db *gorm.DB) {
	if len(c.ContactType) > 25 {
		_ = db.AddError(errors.New("contact type should be less or equal to 25"))
	}
	if len(c.ContactType) == 0 {
		_ = db.AddError(errors.New("contact type should not be empty"))
	}
	if len(c.Description) > 400 {
		_ = db.AddError(errors.New("contact description should not be empty"))
	}
}
