package user

type Pin struct {
	ID      uint   `gorm:"primary_key"`
	UserID  uint   `gorm:"unique_index:idx1"`
	IPNSKey string `gorm:"type:binary(255);unique_index:idx1"`
}
