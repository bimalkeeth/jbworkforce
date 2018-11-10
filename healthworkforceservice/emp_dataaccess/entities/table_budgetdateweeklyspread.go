package entities

import "github.com/jinzhu/gorm"

type TableBudgetDateWeeklySpread struct {
	gorm.Model
	DateWeekId   uint    `gorm:"column:dateweekid"`
	EftProfileId uint    `gorm:"column:eftprofileid"`
	SkillTypeId  uint    `gorm:"column:skilltypeid"`
	WeeklyEft    float64 `gorm:"column:weeklyeft"`
	WeeklyBudget float64 `gorm:"column:weeklybudget"`
}

func (c TableBudgetDateWeeklySpread) TableName() string {
	return "table_budgetdateweeklyspread"
}
