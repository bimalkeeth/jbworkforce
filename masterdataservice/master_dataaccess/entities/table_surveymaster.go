package entities

import "github.com/jinzhu/gorm"

type TableSurveyMaster struct {
	gorm.Model
	SurveyName        string `gorm:"column:surveyname;not_null"`
	SurveyDescription string `gorm:"column:surveydescription;not_null"`
}

func (c TableSurveyMaster) TableName() string {
	return "table_surveymaster"
}
