package entities

import "github.com/jinzhu/gorm"

type TablePatientInfo struct {
	gorm.Model
	URNumber        string `gorm:"column:urnumber"`
	FirstName       string `gorm:"column:firstname"`
	LastName        string `gorm:"column:lastname"`
	IsAcute         bool   `gorm:"column:isacute"`
	IsAggression    bool   `gorm:"column:isaggression"`
	IsSuicidal      bool   `gorm:"column:issuicidal"`
	IsWandering     bool   `gorm:"column:iswandering"`
	IsConfusion     bool   `gorm:"column:isconfusion"`
	IsSelfHarm      bool   `gorm:"column:isselfharm"`
	IsDisInhibited  bool   `gorm:"column:isdisinhibited"`
	IsFalls         bool   `gorm:"column:isfalls"`
	IsHomicidal     bool   `gorm:"column:ishomicidal"`
	IsMedicalAcuity bool   `gorm:"column:ismedicalacuity"`
	IsOther         bool   `gorm:"column:isother"`
	Comments        string `gorm:"column:comments"`
	IsActive        bool   `gorm:"column:isactive"`
}

func (c TablePatientInfo) TableName() string {
	return "table_patientinfo"
}
