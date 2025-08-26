package modules

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `gorm:"size:100;not null"`
	Email string `gorm:"unique;not null"`
	Books []Book
}
