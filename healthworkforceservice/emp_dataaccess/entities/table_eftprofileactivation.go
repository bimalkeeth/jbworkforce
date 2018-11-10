package entities

import "github.com/jinzhu/gorm"

type TableEftProfileActivation struct {
	gorm.Model
	DepartmentId    uint            `gorm:"column:departmentid;not_null"`
	DateWeekId      uint            `gorm:"column:dateweekid;not_null"`
	EftProfileId    uint            `gorm:"column:eftprofileid;not_null"`
	TableEftProfile TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableEftProfileActivation) TableName() string {
	return "table_eftprofileactivation"
}
