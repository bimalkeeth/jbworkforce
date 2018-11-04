package entities

import "github.com/jinzhu/gorm"

type TableQualificationType struct {
	gorm.Model
	QualificationTypeName string `gorm:"column:qualificationtypename;not_null"`
	QualificationTypeAbbr string `gorm:"column:qualificationtypeabbr;not_null"`
}
