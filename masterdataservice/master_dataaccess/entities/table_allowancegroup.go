package entities

import "github.com/jinzhu/gorm"

type TableAllowanceGroup struct {
	gorm.Model
	AllowanceGroupName string `gorm:"column:allowancegroupname;type:varchar(100);not_null"`
	AllowanceGroupAbbr string `gorm:"column:allowancegroupabbr;type:varchar(50);not_null"`
	IsActive           string `gorm:"column:isactive;not_null;default:true"`
	Allowances         []TableAllowance
}

func (c TableAllowanceGroup) TableName() string {
	return "table_allowance"
}
