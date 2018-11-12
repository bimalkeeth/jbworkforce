package entities

import "github.com/jinzhu/gorm"

type TableHealthActivity struct {
	gorm.Model
	Year             int     `gorm:"column:year;not_null"`
	CampusId         uint    `gorm:"column:campusid;not_null"`
	HealthGroupId    uint    `gorm:"column:healthgroupid;not_null"`
	ClinicalStreamId uint    `gorm:"column:clinicalstreamid;not_null"`
	ActivityTypeId   uint    `gorm:"column:activitytypeid;not_null"`
	TotalCount       float64 `gorm:"column:totalcount;default:0.0"`
}

func (c TableHealthActivity) TableName() string {
	return "table_healthactivity"
}
