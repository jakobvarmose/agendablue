package item

import (
	"time"
)

type Item struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID uint
	Data   string `gorm:"type:mediumblob"`
}
