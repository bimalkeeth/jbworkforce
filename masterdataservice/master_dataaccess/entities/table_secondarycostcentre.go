package entities

import "github.com/jinzhu/gorm"

type TableSecondaryCostCentre struct {
	gorm.Model
	SecondaryDepartmentId    uint                     `gorm:"column:secondarydepartmentid;not_null"`
	SecondaryCostCentreName  string                   `gorm:"column:secondarycostcentrename;not_null"`
	SecondaryCostCentreAbbr  string                   `gorm:"column:secondarycostcentreabbr;not_null"`
	IsActive                 bool                     `gorm:"column:isactive;default:true"`
	TableSecondaryDepartment TableSecondaryDepartment `gorm:"foreignkey:secondarydepartmentid"`
}

func (c TableSecondaryCostCentre) TableName() string {
	return "table_secondarycostcentre"
}
