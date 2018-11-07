package entities

import "github.com/jinzhu/gorm"

type TableBudgetData struct {
	gorm.Model
	EftProfileId          uint    `gorm:"column:eftprofileid"`
	EftProfileName        string  `gorm:"column:eftprofilename"`
	StatusId              uint    `gorm:"column:statusid"`
	ProfileStatus         string  `gorm:"column:profilestatus"`
	HospitalId            uint    `gorm:"column:hospitalid"`
	HospitalAbbr          string  `gorm:"column:description"`
	CampusId              uint    `gorm:"column:description"`
	CampusName            string  `gorm:"column:description"`
	DivisionId            uint    `gorm:"column:description"`
	DivisionName          string  `gorm:"column:description"`
	DepartmentId          uint    `gorm:"column:description"`
	DepartmentName        string  `gorm:"column:description"`
	RosterUnitId          int     `gorm:"column:description"`
	RosterUnitName        string  `gorm:"column:description"`
	SecondaryCostCentreId uint    `gorm:"column:description"`
	CostCentreIds         string  `gorm:"column:description"`
	CostCentreNames       string  `gorm:"column:description"`
	HealthCategoryId      uint    `gorm:"column:description"`
	HealthCategoryAbbr    string  `gorm:"column:description"`
	ClinicalStreamId      uint    `gorm:"column:description"`
	ClinicalStreamAbbr    string  `gorm:"column:description"`
	HealthGroupId         uint    `gorm:"column:description"`
	HealthGroupName       string  `gorm:"column:description"`
	GradeCodeId           uint    `gorm:"column:description"`
	GradeCodeAbbr         string  `gorm:"column:description"`
	GradeProfessionId     uint    `gorm:"column:description"`
	GradeProfessionName   string  `gorm:"column:description"`
	AllowanceSkillMixId   uint    `gorm:"column:description"`
	SkillTypeId           uint    `gorm:"column:description"`
	SkillTypeAbbr         string  `gorm:"column:description"`
	StaffTypeId           uint    `gorm:"column:description"`
	StaffTypeName         string  `gorm:"column:description"`
	AllowanceId           uint    `gorm:"column:description"`
	AllowanceAbbr         string  `gorm:"column:description"`
	AllowanceGroupName    string  `gorm:"column:description"`
	Budget                float64 `gorm:"column:description"`
	Eft                   float64 `gorm:"column:description"`
	BedHour               int     `gorm:"column:description"`
	RosterEft             float64 `gorm:"column:description"`
	BudgetBaseEft         float64 `gorm:"column:description"`
	AnnualLeaveEft        float64 `gorm:"column:description"`
	ADOEft                float64 `gorm:"column:description"`
	SickEft               float64 `gorm:"column:description"`
	PGradStudyEft         float64 `gorm:"column:description"`
	ExtraBudgetEft        float64 `gorm:"column:description"`
	BudgetEft             float64 `gorm:"column:description"`
}
