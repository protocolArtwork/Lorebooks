package schema

import (
	"time"

	"gorm.io/gorm"
)

type BookEntry struct {
	gorm.Model

	EntryID       uint  `gorm:"primaryKey autoIncrement"`
	CreatedAt     int64 `gorm:"autoCreateTime index"`
	LastUpdatedAt int64 `gorm:"autoUpdateTime"`

	Filename string
	Author   string

	IsVoid   bool
	VoidedAt int64

	ContentType ContentTypeId
	Content     interface{}
}

func NewEntry(filename string, content interface{}, contentType ContentTypeId) BookEntry {
	return BookEntry{
		Filename: filename,
		Author:   "Unknown Author",
		IsVoid:   false,
		VoidedAt: 0,

		Content:     content,
		ContentType: contentType,
	}
}

func (meta *BookEntry) Retcon() {
	meta.VoidedAt = time.Now().Unix()
	meta.IsVoid = true
}
