package commands

import "github.com/jinzhu/gorm"

type State struct {
	DB     *gorm.DB
	Domain string
}
