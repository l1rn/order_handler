package repositories

import (
	"errors"

	"github.com/l1rn/order-handler/internal/models"
	"gorm.io/gorm"
)

type WorkRepository interface {
	GetAll() ([]models.WorkItem, error)
	Create(w *models.WorkItem) error
	Update(id uint, data map[string]interface{}) error
}

type workRepository struct {
	db *gorm.DB
}

func NewWorkRepository(db *gorm.DB) WorkRepository {
	return &workRepository{db: db}
}

func (r *workRepository) GetAll() ([]models.WorkItem, error) {
	var workItems []models.WorkItem
	err := r.db.Find(&workItems).Error

	if err != nil {
		return nil, errors.New("failed to find work items in db")
	}

	return workItems, err
}

func (r *workRepository) Create(w *models.WorkItem) error {
	return r.db.Create(w).Error
}

func (r *workRepository) Update(
	id uint,
	data map[string]interface{},
) error {
	return r.db.Model(&models.WorkItem{}).Where("id = ?", id).Updates(data).Error
}
