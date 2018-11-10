package entities

import "github.com/jinzhu/gorm"

type TableBudgetData struct {
	gorm.Model
	EftProfileId          uint    `gorm:"column:eftprofileid"`
	EftProfileName        string  `gorm:"column:eftprofilename"`
	StatusId              uint    `gorm:"column:statusid"`
	ProfileStatus         string  `gorm:"column:profilestatus"`
	HospitalId            uint    `gorm:"column:hospitalid"`
	HospitalAbbr          string  `gorm:"column:hospitalabbr"`
	CampusId              uint    `gorm:"column:campusid"`
	CampusName            string  `gorm:"column:campusname"`
	DivisionId            uint    `gorm:"column:divisionid"`
	DivisionName          string  `gorm:"column:divisionname"`
	DepartmentId          uint    `gorm:"column:departmentid"`
	DepartmentName        string  `gorm:"column:departmentname"`
	RosterUnitId          int     `gorm:"column:rosterunitid"`
	RosterUnitName        string  `gorm:"column:rosterunitname"`
	SecondaryCostCentreId uint    `gorm:"column:secondarycostcentreid"`
	CostCentreIds         string  `gorm:"column:costcentreids"`
	CostCentreNames       string  `gorm:"column:costcentrenames"`
	HealthCategoryId      uint    `gorm:"column:healthcategoryid"`
	HealthCategoryAbbr    string  `gorm:"column:healthcategoryabbr"`
	ClinicalStreamId      uint    `gorm:"column:clinicalstreamid"`
	ClinicalStreamAbbr    string  `gorm:"column:clinicalstreamabbr"`
	HealthGroupId         uint    `gorm:"column:healthgroupid"`
	HealthGroupName       string  `gorm:"column:healthgroupname"`
	GradeCodeId           uint    `gorm:"column:gradecodeid"`
	GradeCodeAbbr         string  `gorm:"column:gradecodeabbr"`
	GradeProfessionId     uint    `gorm:"column:gradeprofessionid"`
	GradeProfessionName   string  `gorm:"column:gradeprofessionname"`
	AllowanceSkillMixId   uint    `gorm:"column:allowanceskillmixId"`
	SkillTypeId           uint    `gorm:"column:skilltypeid"`
	SkillTypeAbbr         string  `gorm:"column:skilltypeabbr"`
	StaffTypeId           uint    `gorm:"column:stafftypeid"`
	StaffTypeName         string  `gorm:"column:stafftypename"`
	AllowanceId           uint    `gorm:"column:allowanceid"`
	AllowanceAbbr         string  `gorm:"column:allowanceabbr"`
	AllowanceGroupName    string  `gorm:"column:allowancegroupname"`
	Budget                float64 `gorm:"column:budget;default:0.00"`
	Eft                   float64 `gorm:"column:eft;default:0.00"`
	BedHour               int     `gorm:"column:bedhour;default:0"`
	RosterEft             float64 `gorm:"column:rostereft;default:0.00"`
	BudgetBaseEft         float64 `gorm:"column:budgetbaseeft;default:0.00"`
	AnnualLeaveEft        float64 `gorm:"column:annualleaveeft;default:0.00"`
	ADOEft                float64 `gorm:"column:adoeft;default:0.00"`
	SickEft               float64 `gorm:"column:sickeft;default:0.00"`
	PGradStudyEft         float64 `gorm:"column:pgradstudyeft;default:0.00"`
	ExtraBudgetEft        float64 `gorm:"column:extrabudgeteft;default:0.00"`
	BudgetEft             float64 `gorm:"column:budgeteft;default:0.00"`
}

func (c TableBudgetData) TableName() string {
	return "table_budgetdata"
}
