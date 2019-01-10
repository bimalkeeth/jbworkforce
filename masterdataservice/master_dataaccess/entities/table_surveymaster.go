package entities

import "github.com/jinzhu/gorm"

type TableSurveyMaster struct {
	gorm.Model
	SurveyName        string `gorm:"column:surveyname;type:varchar(200);not_null"`
	SurveyDescription string `gorm:"column:surveydescription;type:varchar(400);not_null"`
}

func (c TableSurveyMaster) TableName() string {
	return "table_surveymaster"
}
