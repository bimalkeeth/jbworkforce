package entities

import "github.com/jinzhu/gorm"

type TableEftProfileGroup struct {
	gorm.Model
	EftProfileGroupName         string `gorm:"column:eftprofilegroupname"`
	DateWeekId                  uint   `gorm:"column:dateweekid;not_null"`
	GradeAwardId                uint   `gorm:"column:gradeawardid;not_null"`
	DateWeekIdCleanup           int    `gorm:"column:dateweekidcleanup"`
	DateWeekIdReport            int    `gorm:"column:dateweekidreport"`
	Description                 string `gorm:"column:description"`
	ProfileStatus               string `gorm:"column:profilestatus"`
	ProfileStage                string `gorm:"column:profilestage"`
	UseShiftLengthAdjustment    bool   `gorm:"column:useshiftlengthadjustment"`
	UseCatStream                bool   `gorm:"column:usecatstream"`
	RosterShiftCalculationClass string `gorm:"column:rostershiftcalclass"`
}

func (c TableEftProfileGroup) TableName() string {
	return "table_eftprofilegroup"
}
