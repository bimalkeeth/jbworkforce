package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableEftProfileAuditAction struct {
	gorm.Model
	EftProfileId    uint            `gorm:"column:eftprofileid;not_null"`
	Action          string          `gorm:"column:action;not_null"`
	ActionTime      time.Time       `gorm:"column:actiontime;not_null"`
	UserId          uint            `gorm:"column:userid;not_null"`
	TableEftProfile TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableEftProfileAuditAction) TableName() string {
	return "table_eftprofileauditaction"
}
