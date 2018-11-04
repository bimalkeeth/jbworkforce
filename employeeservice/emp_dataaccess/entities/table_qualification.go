package entities

import "github.com/jinzhu/gorm"

type TableQualifications struct {
	gorm.Model
	QualificationName      string                 `gorm:"column:qualificationname;not_null"`
	Description            string                 `gorm:"column:description;not_null"`
	QualificationTypeId    uint                   `gorm:"column:qualificationtypeid;not_null"`
	TableQualificationType TableQualificationType `gorm:"foreignkey:qualificationtypeid"`
}
