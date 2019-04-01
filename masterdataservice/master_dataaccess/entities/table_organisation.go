package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableOrganisation struct {
	gorm.Model
	OrganisationName string `gorm:"column:organisationname;type:varchar(128);not_null"`
	OrganisationAbbr string `gorm:"column:organisationabbr;type:varchar(128);not_null"`
	Description      string `gorm:"column:description;type:varchar(400)"`
	IsActive         bool   `gorm:"column:isactive;default:true"`
	EffCalc          int    `gorm:"column:eftcalc;default:0"`
	Hospitals        []*TableHospital
}

func (c TableOrganisation) TableName() string {
	return "table_organisation"
}

func (c TableOrganisation) Validate(db *gorm.DB) {
	if len(c.OrganisationAbbr) > 100 {
		_ = db.AddError(errors.New("organisation abbr length should be less or equal to 100"))
	}
	if c.OrganisationAbbr == "" {
		_ = db.AddError(errors.New("organisation abbr should not be empty"))
	}
	if len(c.OrganisationName) > 300 {
		_ = db.AddError(errors.New("organisation name length should be less or equal to 300"))
	}
	if c.OrganisationName == "" {
		_ = db.AddError(errors.New("organisation name should not be empty"))
	}
	if len(c.Description) > 400 {
		_ = db.AddError(errors.New("description length should be less or equal to 400"))
	}
}
