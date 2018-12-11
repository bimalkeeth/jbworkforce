package entities

import "github.com/jinzhu/gorm"

type TableHospital struct {
	gorm.Model
	OrganisationId    uint              `gorm:"column:organisationid;not_null"`
	HospitalName      string            `gorm:"column:hospitalname;not_null"`
	HospitalAbbr      string            `gorm:"column:hospitalabbr;not_null"`
	HasWorkforce      bool              `gorm:"column:hasworkforce;not_null"`
	HasEftProfiler    bool              `gorm:"column:haseftprofiler;not_null"`
	HasYearlyPlanner  bool              `gorm:"column:hasyearlyplanner;not_null"`
	HasAllocation     bool              `gorm:"column:hasallocation;not_null"`
	IsActive          bool              `gorm:"column:isactive;not_null"`
	Description       string            `gorm:"column:description;not_null"`
	HospitalLogo      string            `gorm:"column:hospitallogo;not_null"`
	TableOrganisation TableOrganisation `gorm:"foreignkey:organisationid"`
	CostCenters       []TableCostCentre `gorm:"foreignkey:hospitalid"`
}

func (c TableHospital) TableName() string {
	return "table_hospital"
}
