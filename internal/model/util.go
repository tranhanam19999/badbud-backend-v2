package model

import (
	"time"

	"github.com/badbud-backend-v2/internal/common/ulidutil"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = ulidutil.NewString()
	}

	return
}
