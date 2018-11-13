package entities

import "github.com/jinzhu/gorm"

type TableRateCalculator struct {
	gorm.Model
	RateCalculatorName string `gorm:"column:ratecalculatorname;not_null"`
	Description        string `gorm:"column:description"`
}

func (c TableRateCalculator) TableName() string {
	return "table_ratecalculator"
}
