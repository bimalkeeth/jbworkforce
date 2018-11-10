package entities

import "github.com/jinzhu/gorm"

type TableCatStreamSkillType struct {
	gorm.Model
	CatStreamId  uint    `gorm:"column:catstreamid;not_null"`
	SkillTypeId  uint    `gorm:"column:skilltypeid;not_null"`
	StaffEft     float64 `gorm:"column:staffeft;default:0.00"`
	Percentage   float64 `gorm:"column:percentage;default:0.00"`
	IsGeneralist bool    `gorm:"column:isgeneralist;default:false"`
}

func (c TableCatStreamSkillType) TableName() string {
	return "table_catstreamskilltype"
}
