package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableGender struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(200);not_null"`
}

func (c TableGender) TableName() string {
	return "table_gender"
}
func (c TableGender) Validate(db *gorm.DB) {
	if len(c.Name) > 200 {
		_ = db.AddError(errors.New("name length should be less or equal to 200"))
	}
}
