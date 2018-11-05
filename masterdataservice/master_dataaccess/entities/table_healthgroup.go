package entities

import "github.com/jinzhu/gorm"

type TableHealthGroup struct {
	gorm.Model
}

func (c TableHealthGroup) TableName() string {
	return "table_healthgroup"
}
