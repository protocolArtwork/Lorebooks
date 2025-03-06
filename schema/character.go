package schema

import (
	"time"

	"gorm.io/gorm"
)

type Character struct {
	gorm.Model

	CharacterId uint `gorm:"index autoIncrement"`

	Name string
	Age  uint8

	ICBornAt time.Time `gorm:"autoCreateTime:milli"`
	ICDiedAt time.Time `gorm:"autoCreateTime:milli"`
}
