package entities

import "github.com/jinzhu/gorm"

type TableNoteType struct {
	gorm.Model
	NoteTypeName        string `gorm:"column:notetypename;type:varchar(200);not_null"`
	NoteTypeDescription string `gorm:"column:notetypedescription;type:varchar(500)"`
}

func (c TableNoteType) TableName() string {
	return "table_notetype"
}
