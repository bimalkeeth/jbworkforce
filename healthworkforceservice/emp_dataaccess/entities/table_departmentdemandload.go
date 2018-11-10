package entities

import "github.com/jinzhu/gorm"

type TableDepartmentDemandLoad struct {
	gorm.Model
	DepartmentId  uint    `gorm:"column:departmentid;not_null"`
	DateWeekId    uint    `gorm:"column:dateWeekid;not_null"`
	Load          float64 `gorm:"column:load;not_null;default:0.00"`
	TargetPercent float64 `gorm:"column:targetpercent;not_null;default:0.00"`
}

func (c TableDepartmentDemandLoad) TableName() string {
	return "table_departmentdemandload"
}
