package entities

import "github.com/jinzhu/gorm"

type TableQualificationType struct {
	gorm.Model
	QualificationTypeName string `gorm:"column:qualificationtypename;type:varchar(200);not_null"`
	QualificationTypeAbbr string `gorm:"column:qualificationtypeabbr;type:varchar(50);not_null"`
	Qualifications        []TableQualifications
}

func (c TableQualificationType) TableName() string {
	return "table_qualificationtype"
}
