package entities

import "github.com/jinzhu/gorm"

type TableCompetencyCompletionStatus struct {
	gorm.Model
	CompetencyCompletionName string `gorm:"column:competencycompletionname;not_null"`
}

func (c TableCompetencyCompletionStatus) TableName() string {
	return "table_competencycompletionstatus"
}
