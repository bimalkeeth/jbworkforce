package entities

import "github.com/jinzhu/gorm"

type TableAllowanceGroup struct {
	gorm.Model
	AllowanceGroupName string           `gorm:"column:allowancegroupname;not_null"`
	AllowanceGroupAbbr string           `gorm:"column:allowancegroupabbr;not_null"`
	IsActive           string           `gorm:"column:isactive;not_null;default:true"`
	Allowances         []TableAllowance `gorm:"foreignkey:allowancegroupid"`
}

func (c TableAllowanceGroup) TableName() string {
	return "table_allowance"
}
