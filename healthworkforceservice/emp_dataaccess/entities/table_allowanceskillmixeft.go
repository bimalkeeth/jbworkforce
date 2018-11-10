package entities

import "github.com/jinzhu/gorm"

type TableAllowanceSkillMixEft struct {
	gorm.Model
	AllowanceSkillMixId uint    `gorm:"column:allowanceskillmixid;not_null"`
	AllowanceId         uint    `gorm:"column:allowanceid;not_null"`
	AllowanceEft        float64 `gorm:"column:allowanceeft;not_null;default:0.00"`
	Type                string  `gorm:"column:type"`
}

func (c TableAllowanceSkillMixEft) TableName() string {
	return "table_allowanceskillmixeft"
}
