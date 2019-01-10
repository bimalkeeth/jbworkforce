package entities

import "github.com/jinzhu/gorm"

type TableSmsStatus struct {
	gorm.Model
	StatusName string `gorm:"column:statusname;type:varchar(200);not_null"`
	SmsGroups  []TableSmsGroup
}

func (c TableSmsStatus) TableName() string {
	return "table_smsstatus"
}
