package entities

import "github.com/jinzhu/gorm"

type TableEftProfileStaffPercent struct {
	gorm.Model
	EftProfileId       uint            `gorm:"column:eftprofileid;not_null"`
	HealthCategoryId   uint            `gorm:"column:healthcategoryid;not_null"`
	ClinicalStreamId   uint            `gorm:"column:clinicalstreamid;not_null"`
	SkillTypeId        uint            `gorm:"column:skilltypeid;not_null"`
	FullTimeEftPercent float64         `gorm:"column:fulltimeeftpercent;default:0.0"`
	PartTimeEftPercent float64         `gorm:"column:parttimeeftpercent;default:0.0"`
	FullTimeStaffCount float64         `gorm:"column:fulltimestaffcount;default:0.0"`
	PartTimeStaffCount float64         `gorm:"column:parttimestaffcount;default:0.0"`
	CntFullTime        float64         `gorm:"column:cntfulltime;default:0.0"`
	CntPartTime        float64         `gorm:"column:cntparttime;default:0.0"`
	EftFullTime        float64         `gorm:"column:eftfulltime;default:0.0"`
	EftPartTime        float64         `gorm:"column:eftparttime;default:0.0"`
	TableEftProfile    TableEftProfile `gorm:"foreignkey:eftprofileid"`
}

func (c TableEftProfileStaffPercent) TableName() string {
	return "table_eftprofilestaffpercent"
}
