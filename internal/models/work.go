package models

import "gorm.io/gorm"

type WorkItem struct {
	gorm.Model
	Name         string
	Description  string
	SubmissionID uint
}

type WorkItemResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Desc         string `json:"desc"`
	SubmissionID uint   `json:"submission_id"`
}

type CreateWorkItemRequest struct {
	Name         *string `json:"name"`
	Desc         *string `json:"desc"`
	SubmissionID uint   `json:"submission_id"`
}
