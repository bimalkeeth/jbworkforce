package entities

import "github.com/jinzhu/gorm"

type TableClinicalStream struct {
	gorm.Model
	HealthGroupId            uint   `gorm:"column:healthgroupid;not_null"`
	ActivityTypeId           uint   `gorm:"column:activitytypeid;not_null"`
	ClinicalStreamName       string `gorm:"column:clinicalstreamname;type:varchar(200);not_null"`
	ClinicalStreamAbbr       string `gorm:"column:clinicalstreamabbr;type:varchar(100);not_null"`
	Description              string `gorm:"column:description"`
	IsClinical               bool   `gorm:"column:isclinical;not_null;default:false"`
	UpdatedUser              string `gorm:"column:updateduser;not_null"`
	HealthGroup              *TableHealthGroup
	ActivityTypes            *TableActivityTypes
	ClinicalStreamAttributes []*TableClinicalStreamAttribute
}

func (c TableClinicalStream) TableName() string {
	return "table_clinicalstream"
}
