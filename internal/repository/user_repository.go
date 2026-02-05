package repository

import (
	"errors"

	"github.com/l1rn/order-handler/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
	GetById(id uint) (*models.User, error)
	Create(*models.User) error
	Delete(*models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetById(id uint) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user).Error
	return user, err
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Delete(user *models.User) error {
	return r.db.Delete(&user).Error
}

func (r *userRepository) UpdatePassword(
	user *models.User,
	oldPassword string,
	newPassword string,
) error {
	// db.Model(user).Update("password", newPassword)
	return nil
}

func (r *userRepository) BeforeDelete(user *models.User) error {
	if user.Role.String() == "admin" {
		return errors.New("admin user not allowed to delete")
	}
	return nil
}
