package entities

import "github.com/jinzhu/gorm"

type TableBusinessRule struct {
	gorm.Model
	BusinessRuleName               string  `gorm:"column:businessrulename"`
	Description                    string  `gorm:"column:description"`
	TotalGradNurse                 int     `gorm:"column:totalgradnurse"`
	TotalPGradNurse                int     `gorm:"column:totalpgradnurse"`
	StudyPercent                   float64 `gorm:"column:studypercent"`
	SickLeavePercent               float64 `gorm:"column:sickleavepercent"`
	OtherEftPercent                float64 `gorm:"column:othereftpercent"`
	OtherTotalEft                  float64 `gorm:"column:othertotaleft"`
	FullTimeAnnualLeave            float64 `gorm:"column:fulltimeannualleave"`
	PartTimeAnnualLeave            float64 `gorm:"column:parttimeannualleave"`
	FullTimeStudyLeave             float64 `gorm:"column:fulltimestudyleave"`
	GradNurseStudyDays             int     `gorm:"column:gradnursestudydays"`
	PGradNurseStudyDays            int     `gorm:"column:pgradnursestudydays"`
	RosterEftForVacancyReplacement float64 `gorm:"column:rostereftforvacancyreplacement"`
	DeleteHoursToEffectiveHours    float64 `gorm:"column:deletehourstoeffectivehours"`
}

func (c TableBusinessRule) TableName() string {
	return "table_businessrule"
}
