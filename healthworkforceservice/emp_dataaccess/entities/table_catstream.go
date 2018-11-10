package entities

import "github.com/jinzhu/gorm"

type TableCatStream struct {
	gorm.Model
	EftProfileId     uint            `gorm:"column:eftprofileid"`
	RosterUnitId     uint            `gorm:"column:rosterunitid"`
	HealthCategoryId uint            `gorm:"column:healthcategoryid"`
	ClinicalStreamId uint            `gorm:"column:clinicalstreamid"`
	RosterEft        float64         `gorm:"column:rostereft"`
	BudgetBaseEft    float64         `gorm:"column:budgetbaseeft"`
	BudgetEft        float64         `gorm:"column:budgeteft"`
	TableRosterUnit  TableRosterUnit `gorm:"foreignkey:rosterunitid"`
	TableEftProfile  TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableCatStream) TableName() string {
	return "table_catstream"
}
