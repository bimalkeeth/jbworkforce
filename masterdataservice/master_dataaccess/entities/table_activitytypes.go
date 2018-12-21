package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableActivityTypes struct {
	gorm.Model
	ActivityTypeName string `gorm:"column:activitytypename;type:varchar(100);not_null"`
	ActivityAbbr     string `gorm:"column:activityabbr;type:varchar(50);not_null"`
	ClinicalStreams  []TableClinicalStream
}

func (c TableActivityTypes) TableName() string {
	return "table_activitytypes"
}

func (c TableActivityTypes) Validate(db *gorm.DB) {
	if len(c.ActivityTypeName) > 100 {
		_ = db.AddError(errors.New("activity type name should be less or equal to 100"))
	}
	if len(c.ActivityAbbr) > 50 {
		_ = db.AddError(errors.New("activity type abbr should be less or equal to 50"))
	}
}
