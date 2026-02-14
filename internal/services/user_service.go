package services

import (
	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/repositories"
)

type UserService interface {
	FindAllUsers() ([]models.UserResponse, error)
	FindById(id uint) (*models.User, error)
	CreateUser(req models.CreateUserRequest) (uint, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
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
			Password: u.Password,
		})
	}
	return response, err
}

func (s *userService) FindById(id uint) (*models.User, error) {
	user, err := s.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return user, err
}

func (s *userService) CreateUser(req models.CreateUserRequest) (uint, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Role:     1,
	}
	err := s.repo.Create(&user)
	return user.ID, err
}
