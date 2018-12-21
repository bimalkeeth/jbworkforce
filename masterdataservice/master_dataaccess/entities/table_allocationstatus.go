package entities

import "github.com/jinzhu/gorm"

type TableAllocationStatus struct {
	gorm.Model
	AllocationStatusName string `gorm:"column:allocationstatusname;type:varchar(200);not_null"`
}
