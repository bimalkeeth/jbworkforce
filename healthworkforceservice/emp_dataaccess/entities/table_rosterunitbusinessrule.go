package entities

import "github.com/jinzhu/gorm"

type TableRosterUnitBusinessRule struct {
	gorm.Model
	TemplateName               string  `gorm:"column:templatename"`
	Description                string  `gorm:"column:description"`
	RosterUnitId               uint    `gorm:"column:rosterunitid;not_null"`
	FullTimeEftPercent         float64 `gorm:"column:fulltimeeftpercent;default:0.0"`
	PartTimeEftPercent         float64 `gorm:"column:parttimeeftpercent;default:0.0"`
	FullTimeStaffCount         float64 `gorm:"column:fulltimestaffcount;default:0.0"`
	PartTimeStaffCount         float64 `gorm:"column:parttimestaffcount;default:0.0"`
	TotalGradNurse             int     `gorm:"column:totalgradnurse;default:0"`
	TotalPGradNurse            int     `gorm:"column:totalpgradnurse;default:0"`
	StudyPercent               float64 `gorm:"column:studypercent;default:0.0"`
	SickLeavePercent           float64 `gorm:"column:sickleavepercent;default:0.0"`
	OtherEftPercent            float64 `gorm:"column:othereftpercent;default:0.0"`
	OtherTotalEft              float64 `gorm:"column:othertotaleft;default:0.0"`
	OtherDescription           string  `gorm:"column:otherdescription"`
	FullTimeAnnualLeave        float64 `gorm:"column:fulltimeannualleave;default:0.0"`
	PartTimeAnnualLeave        float64 `gorm:"column:parttimeannualleave;default:0.0"`
	FullTimeStudyLeave         float64 `gorm:"column:fulltimestudyleave;default:0.0"`
	PartTimeStudyLeave         float64 `gorm:"column:parttimestudyleave;default:0.0"`
	GradNurseStudyDays         int     `gorm:"column:gradnursestudydays;default:0"`
	PGradNurseStudyDays        int     `gorm:"column:pgradnursestudydays;default:0"`
	CasualVacancyReplacement   float64 `gorm:"column:casualvacancyreplacement;default:0.0"`
	ShiftLengthAdjustmentValue float64 `gorm:"column:shiftlengthadjustmentvalue;default:0.0"`
	ExtraBudgetEft             float64 `gorm:"column:extrabudgeteft;default:0.0"`
}

func (c TableRosterUnitBusinessRule) TableName() string {
	return "table_rosterunitbusinessrule"
}
