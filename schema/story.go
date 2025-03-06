package schema

import (
	"time"

	"gorm.io/gorm"
)

type Story struct {
	gorm.Model

	Start time.Time `gorm:"autoCreateTime:milli"`
	End   time.Time `gorm:"autoCreateTime:milli"`
}
