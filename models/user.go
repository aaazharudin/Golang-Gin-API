package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Articles []Article `gorm:"foreignKey:UserID"`
	Username string
	Fullname string
	Email    string
	SocialID string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
}
