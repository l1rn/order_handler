package services

import (
	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/repositories"
)

type SubmissionService interface {
	FindAllUsers() ([]models.SubmissionResponse, error)
	AddWorkItemToSubmission(id uint, wi_id uint) error
	DeleteWorkItemToSubmission(id uint, wi_id uint) error
}

type submissionService struct {
	repo repositories.SubmissionRepository
}

func NewSubmissionService(repo repositories.SubmissionRepository) SubmissionService {
	return &submissionService{repo: repo}
}

func (s *submissionService) FindAllUsers() ([]models.SubmissionResponse, error) {
	var submissions []models.Submission
	submissions, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	var response []models.SubmissionResponse
	for _, s := range submissions {
		var workItems []models.WorkItemResponse
		for _, w := range s.WorkItems {
			workItems = append(workItems, models.WorkItemResponse{
				ID:   w.ID,
				Name: w.Name,
				Desc: w.Description,
			})
		}

		response = append(response, models.SubmissionResponse{
			ID: s.ID,
			User: models.UserResponse{
				ID:       s.User.ID,
				Username: s.User.Username,
				Password: s.User.Password,
				Role:     s.User.Role.String(),
			},
			WorkItems:       workItems,
			SumbmissionDate: s.SubmissionDate,
		})
	}

	return response, nil
}

func (s *submissionService) AddWorkItemToSubmission(id uint, wi_id uint) error {
	err := s.repo.AddWorkItem(id, wi_id)
	if err != nil {
		return err
	}
	return nil
}

func (s *submissionService) DeleteWorkItemToSubmission(id uint, wi_id uint) error {
	err := s.repo.DeleteWorkItem(id, wi_id)
	if err != nil {
		return err
	}
	return nil
}
