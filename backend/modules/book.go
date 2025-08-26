package modules

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `gorm:"size:200;not null"`
	Author string `gorm:"size:100"`
	UserID uint
}
