package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Surname  string
	Username string
	Posts    []Post `gorm:"foreignKey:UserId"`
}
