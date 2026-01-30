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

var roleName = map[UserRole]string {
	RoleUser: "user",
	RoleModerator: "moderator",
	RoleAdmin: "admin",
}

type WorkItem struct {
	gorm.Model
	Name 		string
	Description string
	SubmissionID uint 
}

type Submission struct{
	gorm.Model
	UserID 			uint
	Work			[]WorkItem `gorm:"foreignKey:SubmissionID"`
	SubmissionDate	time.Time 
}

type User struct {
	gorm.Model
	Username 	string		`gorm:"unique"`
	Password 	string 	
	Role		UserRole	`gorm:"default:0"`
	Submission 	[]Submission `gorm:"foreignKey:UserID"`
	CreatedAt 	time.Time
}