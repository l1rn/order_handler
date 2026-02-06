package models

import (
	"time"

	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	UserID         uint
	User           User       `gorm:"foreignKey:UserID"`
	WorkItems      []WorkItem `gorm:"many2many"`
	SubmissionDate time.Time
}

type SubmissionResponse struct {
	User            UserResponse       `json:"user"`
	WorkItems       []WorkItemResponse `json:"work-items"`
	SumbmissionDate time.Time          `json:"date"`
}
