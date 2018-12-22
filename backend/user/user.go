package user

import (
	"time"
)

type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Username string `gorm:"unique"`

	AccessKey  []byte
	ContentKey []byte

	Info      []byte `gorm:"type:blob"`
	Bootstrap []byte `gorm:"type:blob"`
	Content   []byte `gorm:"type:mediumblob"`
}
