package entities

import "github.com/jinzhu/gorm"

type TableSurveyQuestions struct {
	gorm.Model
	SurveyId          uint              `gorm:"column:surveyid;not_null"`
	Question          string            `gorm:"column:question;not_null"`
	Weight            float64           `gorm:"column:weight;not_null"`
	TableSurveyMaster TableSurveyMaster `gorm:"foreignkey:surveyid"`
}

func (c TableSurveyQuestions) TableName() string {
	return "table_surveyquestions"
}
