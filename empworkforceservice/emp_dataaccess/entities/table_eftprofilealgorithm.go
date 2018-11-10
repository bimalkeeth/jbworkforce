package entities

import "github.com/jinzhu/gorm"

type TableEftProfileAlgorithm struct {
	gorm.Model
	AlgorithmType  string `gorm:"column:algorithmtype"`
	AlgorithmName  string `gorm:"column:algorithmname"`
	Description    string `gorm:"column:description"`
	AlgorithmClass string `gorm:"column:algorithmclass"`
	IsActive       bool   `gorm:"column:isactive;default:false"`
}

func (c TableEftProfileAlgorithm) TableName() string {
	return "table_eftprofilealgorithm"
}
