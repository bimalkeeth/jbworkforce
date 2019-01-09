package entities

import "github.com/jinzhu/gorm"

type TableRateCalculator struct {
	gorm.Model
	RateCalculatorName string `gorm:"column:ratecalculatorname;type:varchar(200);not_null"`
	Description        string `gorm:"column:description;type:varchar(400)"`
}

func (c TableRateCalculator) TableName() string {
	return "table_ratecalculator"
}
