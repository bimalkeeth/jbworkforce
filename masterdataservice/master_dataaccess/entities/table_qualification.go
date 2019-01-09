package entities

import "github.com/jinzhu/gorm"

type TableQualifications struct {
	gorm.Model
	QualificationName   string `gorm:"column:qualificationname;type:varchar(300);not_null"`
	Description         string `gorm:"column:description;type:varchar(400);not_null"`
	QualificationTypeId uint   `gorm:"column:qualificationtypeid;not_null"`
	QualificationType   TableQualificationType
}

func (c TableQualifications) TableName() string {
	return "table_qualifications"
}
