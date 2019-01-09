package entities

import "github.com/jinzhu/gorm"

type TableHospital struct {
	gorm.Model
	OrganisationId   uint   `gorm:"column:organisationid;not_null"`
	HospitalName     string `gorm:"column:hospitalname;type:varchar(300);not_null"`
	HospitalAbbr     string `gorm:"column:hospitalabbr;type:varchar(100);not_null"`
	HasWorkforce     bool   `gorm:"column:hasworkforce;not_null"`
	HasEftProfiler   bool   `gorm:"column:haseftprofiler;not_null"`
	HasYearlyPlanner bool   `gorm:"column:hasyearlyplanner;not_null"`
	HasAllocation    bool   `gorm:"column:hasallocation;not_null"`
	IsActive         bool   `gorm:"column:isactive;not_null"`
	Description      string `gorm:"column:description;type:varchar(400);not_null"`
	HospitalLogo     string `gorm:"column:hospitallogo;type:varchar(200);not_null"`
	Organisation     TableOrganisation
	CostCenters      []TableCostCentre
}

func (c TableHospital) TableName() string {
	return "table_hospital"
}
