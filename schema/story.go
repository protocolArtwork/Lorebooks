package schema

import (
	"time"

	"gorm.io/gorm"
)

type StoryChapter struct {
	gorm.Model

	InnerHTML string

	Start time.Time `gorm:"autoCreateTime:milli"`
	End   time.Time `gorm:"autoCreateTime:milli"`
}
