package user

import "time"

type Login struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	UserID uint

	LoginKey []byte
}
