package models

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRole int

const (
	RoleUser UserRole = iota
	RoleModerator
	RoleAdmin
)

var roleName = map[UserRole]string{
	RoleUser:      "user",
	RoleModerator: "moder",
	RoleAdmin:     "admin",
}

type User struct {
	gorm.Model
	Username   string       `gorm:"unique"`
	Password   string       `gorm:"not null"`
	Role       UserRole     `gorm:"default:0"`
	Submission []Submission `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type UserPasswordRequest struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type UserResponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (r UserRole) String() string {
	if name, ok := roleName[r]; ok {
		return name
	}
	return "unknown"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password for user")
	}
	u.Password = string(hashedPassword)
	return nil
}
