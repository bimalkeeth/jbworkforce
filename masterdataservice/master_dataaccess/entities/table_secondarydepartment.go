package entities

import "github.com/jinzhu/gorm"

type TableSecondaryDepartment struct {
	gorm.Model
	SecondaryDivisionId     uint   `gorm:"column:secondarydivisionid;not_null"`
	SecondaryDepartmentName string `gorm:"column:secondarydepartmentname;type:varchar(200);not_null"`
	SecondaryDepartmentAbbr string `gorm:"column:secondarydepartmentabbr;type:varchar(50);not_null"`
	IsActive                bool   `gorm:"column:isactive;default:true"`
}

func (c TableSecondaryDepartment) TableName() string {
	return "table_secondarydepartment"
}
