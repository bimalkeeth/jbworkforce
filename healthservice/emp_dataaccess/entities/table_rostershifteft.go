package entities

import "github.com/jinzhu/gorm"

type TableRosterShiftEft struct {
	gorm.Model
	RosterShiftId    uint             `gorm:"column:rostershiftid;not_null"`
	RosterHourBF     float64          `gorm:"column:rosterhourbf;default:0.0"`
	RosterHourNBF    float64          `gorm:"column:rosterhournbf;default:0.0"`
	RosterEftBF      float64          `gorm:"column:rostereftbf;default:0.0"`
	RosterEftNBF     float64          `gorm:"column:rostereftnbf;default:0.0"`
	RosterEft        float64          `gorm:"column:rostereft;default:0.0"`
	AnnualLeaveEft   float64          `gorm:"column:annualleaveeft;default:0.0"`
	ADOEft           float64          `gorm:"column:adoeft;default:0.0"`
	BudgetBaseEftBF  float64          `gorm:"column:budgetbaseeftbf;default:0.0"`
	BudgetBaseEft    float64          `gorm:"column:budgetbaseeft;default:0.0"`
	SickEft          float64          `gorm:"column:sickeft;default:0.0"`
	StudyEft         float64          `gorm:"column:studyeft;default:0.0"`
	PGradStudyEft    float64          `gorm:"column:pgradstudyeft;default:0.0"`
	BudgetEftBF      float64          `gorm:"column:budgeteftbf;default:0.0"`
	BudgetEft        float64          `gorm:"column:budgeteft;default:0.0"`
	TableRosterShift TableRosterShift `gorm:"foreignkey:rostershiftid"`
}

func (c TableRosterShiftEft) TableName() string {
	return "table_rostershifteft"
}
