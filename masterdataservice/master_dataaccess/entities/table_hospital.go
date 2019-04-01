package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableHospital struct {
	gorm.Model
	OrganisationId   uint   `gorm:"column:organisationid;not_null"`
	HospitalName     string `gorm:"column:hospitalname;type:varchar(300);not_null"`
	HospitalAbbr     string `gorm:"column:hospitalabbr;type:varchar(100);not_null"`
	HasWorkforce     bool   `gorm:"column:hasworkforce;not_null;default:false"`
	HasEftProfiler   bool   `gorm:"column:haseftprofiler;not_null;default:false"`
	HasYearlyPlanner bool   `gorm:"column:hasyearlyplanner;not_null;default:false"`
	HasAllocation    bool   `gorm:"column:hasallocation;not_null;default:false"`
	IsActive         bool   `gorm:"column:isactive;not_null;default:true"`
	Description      string `gorm:"column:description;type:varchar(400);not_null"`
	HospitalLogo     string `gorm:"column:hospitallogo;type:varchar(200);not_null"`
	Organisation     *TableOrganisation
	CostCenters      []*TableCostCentre
}

func (c TableHospital) TableName() string {
	return "table_hospital"
}

func (c TableHospital) Validate(db *gorm.DB) {
	if len(c.HospitalAbbr) > 100 {
		_ = db.AddError(errors.New("hospital abbr length should be less or equal to 100"))
	}
	if c.HospitalAbbr == "" {
		_ = db.AddError(errors.New("hospital abbr should not be empty"))
	}
	if len(c.HospitalName) > 300 {
		_ = db.AddError(errors.New("hospital name length should be less or equal to 300"))
	}
	if c.HospitalName == "" {
		_ = db.AddError(errors.New("costcentre name should not be empty"))
	}
	if len(c.Description) > 400 {
		_ = db.AddError(errors.New("description length should be less or equal to 400"))
	}
}
