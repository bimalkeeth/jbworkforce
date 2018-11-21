package entities

import "github.com/jinzhu/gorm"

type TableWorkforceRequestMember struct {
	gorm.Model
	WorkforceRequestId    uint                  `gorm:"column:workforcerequestid;not_null"`
	UserId                uint                  `gorm:"column:userid;not_null"`
	TableWorkforceRequest TableWorkforceRequest `gorm:"foreignkey:workforcerequestid"`
}

func (c TableWorkforceRequestMember) TableName() string {
	return "table_workforcerequestmember"
}
