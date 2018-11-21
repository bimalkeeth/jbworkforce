package entities

import "github.com/jinzhu/gorm"

type TableEftProfileActivity struct {
	gorm.Model
	ActivityName     string          `gorm:"column:activityname;not_null"`
	EftProfileId     uint            `gorm:"column:eftprofileid;not_null"`
	ActivityTypeId   uint            `gorm:"column:activitytypeid;not_null"`
	ClinicalStreamId uint            `gorm:"column:clinicalstreamid;not_null"`
	HealthCategoryId uint            `gorm:"column:healthcategoryid;not_null"`
	TotalCount       float64         `gorm:"column:totalcount;default:0.00"`
	Description      string          `gorm:"column:description"`
	TableEftProfile  TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableEftProfileActivity) TableName() string {
	return "table_profileactivity"
}
