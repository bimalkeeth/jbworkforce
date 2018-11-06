package entities

import "github.com/jinzhu/gorm"

type TableAllowanceSkillMix struct {
	gorm.Model
	EftProfileId     uint            `gorm:"column:eftprofileid;not_null"`
	HealthCategoryId uint            `gorm:"column:healthcategoryid;not_null"`
	ClinicalStreamId uint            `gorm:"column:clinicalstreamid;not_null"`
	SkillTypeId      uint            `gorm:"column:skilltypeid;not_null"`
	GradeCodeId      uint            `gorm:"column:gradecodeid;not_null"`
	StaffTypeId      uint            `gorm:"column:stafftypeid;not_null"`
	RosterEft        float64         `gorm:"column:rostereft;not_null;default:0.00"`
	AnnualLeaveEft   float64         `gorm:"column:annualleaveeft;not_null;default:0.00"`
	ADOEft           float64         `gorm:"column:adoeft;not_null;default:0.00"`
	BudgetBaseEft    float64         `gorm:"column:budgetbaseeft;not_null;default:0.00"`
	SickEft          float64         `gorm:"column:sickeft;not_null;default:0.00"`
	StudyEft         float64         `gorm:"column:studyeft;not_null;default:0.00"`
	PGradStudyEft    float64         `gorm:"column:pgradstudyeft;not_null;default:0.00"`
	ExtraBudgetEft   float64         `gorm:"column:extrabudgeteft;not_null;default:0.00"`
	Eft              float64         `gorm:"column:eft;not_null;default:0.00"`
	ShiftWorkerEft   float64         `gorm:"column:shiftworkereft;not_null;default:0.00"`
	GroupId          uint            `gorm:"column:groupid;not_null"`
	IsGenerated      bool            `gorm:"column:isgenerated;not_null;default:false"`
	TableEftProfile  TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableAllowanceSkillMix) TableName() string {
	return "table_allowanceskillmix"
}
