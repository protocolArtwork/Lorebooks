package schema

import (
	"time"

	"gorm.io/gorm"
)

type StoryEvent struct {
	gorm.Model

	Metadata // include metadata

	StoryEventID uint64 `gorm:"primary_key AUTO_INCREMENT"`

	Start time.Time `gorm:"autoCreateTime:milli"`
	End   time.Time `gorm:"autoCreateTime:milli"`
}
