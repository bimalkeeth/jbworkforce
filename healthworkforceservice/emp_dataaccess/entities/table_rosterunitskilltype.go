package entities

import "github.com/jinzhu/gorm"

type TableRosterUnitSkillType struct {
	gorm.Model
	RosterUnitId    uint            `gorm:"column:rosterunitid;not_null"`
	SkillTypeId     uint            `gorm:"column:skilltypeid;not_null"`
	Active          bool            `gorm:"column:active;default:true"`
	Hour            float64         `gorm:"column:hour;default:0.0"`
	Eft             float64         `gorm:"column:eft;default:0.0"`
	IsGeneralist    bool            `gorm:"column:isgeneralist;default:false"`
	TableRosterUnit TableRosterUnit `gorm:"foreignkey:rosterunitid"`
}

func (c TableRosterUnitSkillType) TableName() string {
	return "table_rosterunitskilltype"
}
