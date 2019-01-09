package entities

import "github.com/jinzhu/gorm"

type TableHealthCategory struct {
	gorm.Model
	HealthCategoryName string `gorm:"column:healthcategoryname;type:varchar(200);not_null"`
	HealthCategoryAbbr string `gorm:"column:healthcategoryabbr;type:varchar(100);not_null"`
	Description        string `gorm:"column:description;type:varchar(300);not_null"`
	IsClinical         bool   `gorm:"column:isclinical;default:false"`
	UserId             uint   `gorm:"column:userid;not_null"`
}

func (c TableHealthCategory) TableName() string {
	return "table_healthcategory"
}
