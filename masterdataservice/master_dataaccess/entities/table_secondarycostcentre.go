package entities

import "github.com/jinzhu/gorm"

type TableSecondaryCostCentre struct {
	gorm.Model
	SecondaryDepartmentId   uint   `gorm:"column:secondarydepartmentid;not_null"`
	SecondaryCostCentreName string `gorm:"column:secondarycostcentrename;type:varchar(200);not_null"`
	SecondaryCostCentreAbbr string `gorm:"column:secondarycostcentreabbr;type:varchar(50);not_null"`
	IsActive                bool   `gorm:"column:isactive;default:true"`
	SecondaryDepartment     TableSecondaryDepartment
}

func (c TableSecondaryCostCentre) TableName() string {
	return "table_secondarycostcentre"
}
