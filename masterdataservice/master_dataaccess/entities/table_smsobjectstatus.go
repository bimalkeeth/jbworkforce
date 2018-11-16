package entities

import "github.com/jinzhu/gorm"

type TableSmsObjectStatus struct {
	gorm.Model
	ObjectType string `gorm:"column:objecttype;not_null"`
	ObjectId   int    `gorm:"column:objectid;not_null"`
}

func (c TableSmsObjectStatus) TableName() string {
	return "table_smsobjectstatus"
}
