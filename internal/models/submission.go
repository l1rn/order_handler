package models

import (
	"time"

	"gorm.io/gorm"
)

type Submission struct {
	gorm.Model
	UserID         uint
	Work           []WorkItem `gorm:"foreignKey:SubmissionID"`
	SubmissionDate time.Time
}
