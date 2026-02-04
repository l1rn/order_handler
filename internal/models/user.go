package models

import (
	"time"

	"gorm.io/gorm"
)

type UserRole int

const (
	RoleUser UserRole = iota
	RoleModerator
	RoleAdmin
)

type User struct {
	gorm.Model
	Username   string `gorm:"unique"`
	Password   string
	Role       UserRole     `gorm:"default:0"`
	Submission []Submission `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}