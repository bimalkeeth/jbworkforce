package entities

import "github.com/jinzhu/gorm"

type TableTemplateSpread struct {
	gorm.Model
	TemplateSpreadName string `gorm:"column:templatespreadname;not_null"`
	Description        string `gorm:"column:description"`
}

func (c TableTemplateSpread) TableName() string {
	return "table_templatespread"
}
