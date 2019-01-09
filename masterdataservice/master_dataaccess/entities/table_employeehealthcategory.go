package entities

import "github.com/jinzhu/gorm"

type TableEmployeeHealthCategory struct {
	gorm.Model
	EmployeeId       uint `gorm:"column:employeeid;not_null"`
	HealthCategoryId uint `gorm:"column:healthcategoryid;not_null"`
	Employee         TableEmployee
	HealthCategory   TableHealthCategory
}

func (c TableEmployeeHealthCategory) TableName() string {
	return "table_employeehealthcategory"
}
