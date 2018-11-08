package entities

import "github.com/jinzhu/gorm"

type TableAllocationStatus struct {
	gorm.Model
	AllocationStatusName string `gorm:"column:allocationstatusname;not_null;size:128"`
}
