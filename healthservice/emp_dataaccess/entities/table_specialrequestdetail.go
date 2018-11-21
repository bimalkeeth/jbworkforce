package entities

import "github.com/jinzhu/gorm"

type TableSpecialRequestDetail struct {
	gorm.Model
	SpecialRequestId   uint   `gorm:"column:specialrequestid;not_null"`
	PatientId          uint   `gorm:"column:patientid;not_null"`
	IsAcute            bool   `gorm:"column:isacute;default:false"`
	IsAggression       bool   `gorm:"column:isaggression;default:false"`
	IsSuicidal         bool   `gorm:"column:issuicidal;default:false"`
	IsWandering        bool   `gorm:"column:iswandering;default:false"`
	IsConfusion        bool   `gorm:"column:isconfusion;default:false"`
	IsSelfHarm         bool   `gorm:"column:isselfharm;default:false"`
	IsDisInhibited     bool   `gorm:"column:isdisinhibited;default:false"`
	IsFalls            bool   `gorm:"column:isfalls;default:false"`
	IsHomicidal        bool   `gorm:"column:ishomicidal;default:false"`
	IsMedicalAcuity    bool   `gorm:"column:ismedicalacuity;default:false"`
	IsOther            bool   `gorm:"column:isother;default:false"`
	Comments           string `gorm:"column:comments"`
	IsSpecialApprove   bool   `gorm:"column:isspecialapprove;default:false"`
	ApproveSpecialName string `gorm:"column:approvespecialname"`
	NurseArmDistance   string `gorm:"column:nursearmdistance"`
}

func (c TableSpecialRequestDetail) TableName() string {
	return "table_specialrequestdetail"
}
