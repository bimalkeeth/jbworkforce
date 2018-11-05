package entities

import "github.com/jinzhu/gorm"

type TableAllowance struct {
	gorm.Model
	AllowanceName       string              `gorm:"column:allowancename;not_null"`
	AllowanceAbbr       string              `gorm:"column:allowanceabbr;not_null"`
	Description         string              `gorm:"column:description"`
	SortOrder           string              `gorm:"column:sortorder"`
	AllowanceGroupId    uint                `gorm:"column:allowancegroupid;not_null"`
	TableAllowanceGroup TableAllowanceGroup `gorm:"foreignkey:allowancegroupid"`
}

func (c TableAllowance) TableName() string {
	return "table_allowance"
}
