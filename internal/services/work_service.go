package services

import (
	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/repositories"
)

type WorkService interface {
	FindAllWorkItems() ([]models.WorkItemResponse, error)
	CreateWorkItem(req models.CreateWorkItemRequest) error
	UpdateWorkItem(id uint, req models.CreateWorkItemRequest) error
}

type workService struct {
	repo repositories.WorkRepository
}

func NewWorkService(repo repositories.WorkRepository) WorkService {
	return &workService{repo: repo}
}

func (s *workService) FindAllWorkItems() ([]models.WorkItemResponse, error) {
	workItems, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	var response []models.WorkItemResponse

	for _, w := range workItems {
		response = append(response, models.WorkItemResponse{
			ID:           w.ID,
			Name:         w.Name,
			Desc:         w.Description,
		})
	}
	return response, err
}

func (s *workService) CreateWorkItem(req models.CreateWorkItemRequest) error {
	workItem := models.WorkItem{
		Name:         *req.Name,
		Description:  *req.Desc,
	}

	return s.repo.Create(&workItem)
}

func (s *workService) UpdateWorkItem(id uint, req models.CreateWorkItemRequest) error {
	updatedData := make(map[string]interface{})

	if req.Name != nil {
		updatedData["name"] = *req.Name
	}

	if req.Desc != nil {
		updatedData["description"] = *req.Desc
	}

	return s.repo.Update(id, updatedData)
}
