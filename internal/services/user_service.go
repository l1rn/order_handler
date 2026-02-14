package services

import (
	"errors"

	"github.com/l1rn/order-handler/internal/models"
	"github.com/l1rn/order-handler/internal/repositories"
	"gorm.io/gorm"
)

type UserService interface {
	FindAllUsers() ([]models.UserResponse, error)
	FindById(id uint) (*models.User, error)
	FindByUsername(username string) (*models.UserResponse, error)
	CreateUser(req models.UserRequest) (uint, error)
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

func (s *userService) FindByUsername(username string) (*models.UserResponse, error) {
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, nil
		}
		return nil, err
	}

	return &models.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role.String(),
	}, err
}

func (s *userService) CreateUser(req models.UserRequest) (uint, error) {
	user := models.User{
		Username: req.Username,
		Password: req.Password,
		Role:     1,
	}
	err := s.repo.Create(&user)
	return user.ID, err
}
