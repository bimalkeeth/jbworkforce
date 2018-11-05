package entities

import "github.com/jinzhu/gorm"

type TableCostCentre struct {
	gorm.Model
}

func (c TableCostCentre) TableName() string {
	return "table_costcentre"
}
