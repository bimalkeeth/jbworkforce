package entities

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type TableDepartment struct {
	gorm.Model
	DepartmentName string `gorm:"column:departmentname;type:varchar(200);not_null"`
	DepartmentAbbr string `gorm:"column:departmentabbr;type:varchar(50);not_null"`
	Description    string `gorm:"column:description;type:varchar(250);"`
	Definition     string `gorm:"column:definition;type:varchar(250);"`
	IsActive       bool   `gorm:"column:isactive;not_null;default:true"`
	DivisionId     uint   `gorm:"column:divisionid;not_null"`
	Division       *TableDivision
	SubDepartments []*TableSubDepartment
}

func (c TableDepartment) TableName() string {
	return "table_department"
}
func (c TableDepartment) Validate(db *gorm.DB) {

	if len(c.DepartmentName) > 200 {
		_ = db.AddError(errors.New("department name length should be less or equal to 128"))
	}
	if len(c.DepartmentAbbr) > 50 {
		_ = db.AddError(errors.New("department abbr length should be less or equal to 50"))
	}
	if len(c.Description) > 250 {
		_ = db.AddError(errors.New("description length should be less or equal to 250"))
	}
	if len(c.Definition) > 250 {
		_ = db.AddError(errors.New("definition length should be less or equal to 250"))
	}
}
