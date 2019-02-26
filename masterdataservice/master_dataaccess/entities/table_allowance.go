package entities

import "github.com/jinzhu/gorm"

type TableAllowance struct {
	gorm.Model
	AllowanceName    string `gorm:"column:allowancename;type:varchar(200);not_null"`
	AllowanceAbbr    string `gorm:"column:allowanceabbr;type:varchar(50);not_null"`
	Description      string `gorm:"column:description;type:varchar(400)"`
	SortOrder        string `gorm:"column:sortorder;default:0"`
	AllowanceGroupId uint   `gorm:"column:allowancegroupid;not_null"`
	AllowanceGroup   *TableAllowanceGroup
}

func (c TableAllowance) TableName() string {
	return "table_allowance"
}
