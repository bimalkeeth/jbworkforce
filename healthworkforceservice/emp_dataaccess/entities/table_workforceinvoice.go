package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableWorkforceInvoice struct {
	gorm.Model
	WorkforceRequestId    uint                  `gorm:"column:workforcerequestid;not_null"`
	AgencyId              uint                  `gorm:"column:agencyid;not_null"`
	Expense               float64               `gorm:"column:expense;default:0.0"`
	InvoiceAmount         float64               `gorm:"column:invoiceamount;default:0.0"`
	InvoiceDate           time.Time             `gorm:"column:templatespreadname"`
	InvoiceNumber         string                `gorm:"column:invoicenumber"`
	IsBilled              bool                  `gorm:"column:isbilled;default:false"`
	IsSlipped             bool                  `gorm:"column:isslipped;default:false"`
	IsApproved            bool                  `gorm:"column:isapproved;default:false"`
	IsPaid                bool                  `gorm:"column:ispaid;default:false"`
	PaidDate              time.Time             `gorm:"column:paiddate;not_null"`
	PaymentName           string                `gorm:"column:paymentname;not_null"`
	PaymentGradeCode      string                `gorm:"column:paymentgradecode"`
	PaymentStaffType      string                `gorm:"column:paymentstafftype"`
	PaymentAgencyName     string                `gorm:"column:paymentagencyname"`
	PaymentAgencyCode     string                `gorm:"column:paymentagencycode"`
	PaymentStartDate      time.Time             `gorm:"column:paymentstartdate"`
	PaymentStartTime      time.Time             `gorm:"column:paymentstarttime"`
	PaymentFinishDate     time.Time             `gorm:"column:paymentfinishdate"`
	PaymentFinishTime     time.Time             `gorm:"column:paymentfinishtime"`
	PaymentDepartment     string                `gorm:"column:paymentdepartment"`
	PaymentCostCentre     string                `gorm:"column:paymentcostcentre"`
	PaymentHourWorked     float64               `gorm:"column:paymenthourworked;default:0.0"`
	PaymentHourDayOne     float64               `gorm:"column:paymenthourdayone;default:0.0"`
	PaymentHourDayTwo     float64               `gorm:"column:paymenthourdaytwo;default:0.0"`
	HolidaysInfo          string                `gorm:"column:holidaysinfo"`
	CostDayOne            float64               `gorm:"column:costdayone;default:0.0"`
	CostDayOneBaseRate    float64               `gorm:"column:costdayonebaserate;default:0.0"`
	CostDayOneRate        float64               `gorm:"column:costdayonerate;default:0.0"`
	CostDayTwo            float64               `gorm:"column:costdaytwo;default:0.0"`
	CostDayTwoBaseRate    float64               `gorm:"column:costdaytwobaserate;default:0.0"`
	CostDayTwoRate        float64               `gorm:"column:costdaytworate;default:0.0"`
	CostStaffTypeRate     float64               `gorm:"column:coststafftyperate;default:0.0"`
	CostShiftAllowance    float64               `gorm:"column:costshiftallowance;default:0.0"`
	CostUniform           float64               `gorm:"column:costuniform;default:0.0"`
	CostQualification     float64               `gorm:"column:costqualification;default:0.0"`
	CostTotal             float64               `gorm:"column:costtotal;default:0.0"`
	Info                  string                `gorm:"column:info"`
	TableWorkforceRequest TableWorkforceRequest `gorm:"foreignkey:workforcerequestid"`
}

func (c TableWorkforceInvoice) TableName() string {
	return "table_workforceinvoice"
}
