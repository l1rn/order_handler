package models

import "gorm.io/gorm"

type WorkItem struct {
	gorm.Model
	Name         string
	Description  string
	SubmissionID uint
}