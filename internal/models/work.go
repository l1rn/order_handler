package models

import "gorm.io/gorm"

type WorkItem struct {
	gorm.Model
	Name        string       `json:"name"`
	Description string       `json:"desc"`
	Submission  []Submission `gorm:"many2many:submission_works"`
}

type WorkItemResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type CreateWorkItemRequest struct {
	Name *string `json:"name"`
	Desc *string `json:"desc"`
}
