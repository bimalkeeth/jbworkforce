package entities

import "github.com/jinzhu/gorm"

type TableSmsStatus struct {
	gorm.Model
	StatusName string `gorm:"column:statusname;not_null"`
}

func (c TableSmsStatus) TableName() string {
	return "table_smsstatus"
}
