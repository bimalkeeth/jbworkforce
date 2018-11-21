package entities

import "github.com/jinzhu/gorm"

type TableEftProfileAllowanceDetail struct {
	gorm.Model
	EftProfileAllowanceId uint `gorm:"column:eftprofileallowanceid;not_null"`
	ClinicalStreamId      uint `gorm:"column:clinicalstreamId;not_null"`
	HealthCategoryId      uint `gorm:"column:healthcategoryid;not_null"`
	RosterUnitId          uint `gorm:"column:rosterunitid;not_null"`
	StaffTypeId           uint `gorm:"column:stafftypeid;not_null"`
}

func (c TableEftProfileAllowanceDetail) TableName() string {
	return "table_eftprofileallowancedetail"
}
