package entities

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TableSmsText struct {
	gorm.Model
	SmsGroupId       uint           `gorm:"column:smsgroupid;not_null"`
	MobileNumber     string         `gorm:"column:mobilenumber;not_null"`
	SentText         string         `gorm:"column:senttext;not_null"`
	SentDateTime     time.Time      `gorm:"column:sentdatetime;not_null"`
	ReplyText        string         `gorm:"column:replytext;not_null"`
	ReplyDateTime    time.Time      `gorm:"column:replydatetime"`
	ConfirmationCode string         `gorm:"column:confirmationcode"`
	SmsStatusId      uint           `gorm:"column:smsstatusid;not_null"`
	Description      string         `gorm:"column:description"`
	TableSmsStatus   TableSmsStatus `gorm:"foreignkey:smsstatusid"`
}

func (c TableSmsText) TableName() string {
	return "table_smstext"
}
