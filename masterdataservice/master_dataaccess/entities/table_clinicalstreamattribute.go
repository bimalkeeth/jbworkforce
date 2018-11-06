package entities

import "github.com/jinzhu/gorm"

type TableClinicalStreamAttribute struct {
	gorm.Model
	IsNewService        bool                `gorm:"column:isnewservice;not_null;default:false"`
	IsGrowthService     bool                `gorm:"column:isgrowthservice;not_null;default:false"`
	IsNewGrowthService  bool                `gorm:"column:isnewgrowthservice;not_null;default:false"`
	ClinicalStreamId    uint                `gorm:"column:clinicalstreamid;not_null"`
	TableClinicalStream TableClinicalStream `gorm:"foreignkey:clinicalstreamid"`
}

func (c TableClinicalStreamAttribute) TableName() string {
	return "table_clinicalstreamattribute"
}
