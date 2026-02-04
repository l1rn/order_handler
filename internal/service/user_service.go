package service

import (
	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/repository"
)

type UserService interface {
	FindAllUsers() ([]models.UserResponse, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) FindAllUsers() ([]models.UserResponse, error) {
	users, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	var response []models.UserResponse
	for _, u := range users {
		response = append(response, models.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Role:     u.Role.String(),
		})
	}
	return response, err
}
