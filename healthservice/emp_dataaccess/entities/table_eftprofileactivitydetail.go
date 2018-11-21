package entities

import "github.com/jinzhu/gorm"

type TableEftProfileActivityDetail struct {
	gorm.Model
	EftProfileActivityId    uint                    `gorm:"column:eftprofileactivityid;not_null"`
	ShiftTypeId             uint                    `gorm:"column:shifttypeid;not_null"`
	GroupId                 uint                    `gorm:"column:groupid;default:0"`
	ActivityLength          float64                 `gorm:"column:activitylength;not_null;default:0.0"`
	Monday                  float64                 `gorm:"column:monday;not_null;0.0"`
	Tuesday                 float64                 `gorm:"column:tuesday;not_null;0.0"`
	Wednesday               float64                 `gorm:"column:wednesday;not_null;0.0"`
	Thursday                float64                 `gorm:"column:thursday;not_null;0.0"`
	Friday                  float64                 `gorm:"column:friday;not_null;0.0"`
	Saturday                float64                 `gorm:"column:saturday;not_null;0.0"`
	Sunday                  float64                 `gorm:"column:sunday;not_null;0.0"`
	Weekly                  float64                 `gorm:"column:weekly;not_null;0.0"`
	TableEftProfileActivity TableEftProfileActivity `gorm:"foreignkey:eftprofileactivityid"`
}

func (c TableEftProfileActivityDetail) TableName() string {
	return "table_eftprofileactivitydetail"
}
