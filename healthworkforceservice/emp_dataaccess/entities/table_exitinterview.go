package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableExitInterview struct {
	gorm.Model
	CompletedSurveyId uint          `gorm:"column:completedsurveyid;not_null"`
	ExitDate          time.Time     `gorm:"column:exitdate;not_null"`
	InterviewDate     time.Time     `gorm:"column:interviewdate;not_null"`
	EmployeeId        uint          `gorm:"column:employeeid;not_null"`
	TableEmployee     TableEmployee `gorm:"foreignkey:employeeid"`
}

func (c TableExitInterview) TableName() string {
	return "table_exitinterview"
}
