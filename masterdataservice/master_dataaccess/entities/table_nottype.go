package entities

import "github.com/jinzhu/gorm"

type TableNoteType struct {
	gorm.Model
	NoteTypeName        string `gorm:"column:notetypename;not_null"`
	NoteTypeDescription string `gorm:"column:notetypedescription"`
}

func (c TableNoteType) TableName() string {
	return "table_notetype"
}
