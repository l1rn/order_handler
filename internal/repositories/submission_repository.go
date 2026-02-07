package repositories

import (
	"github.com/l1rn/order-handler/internal/models"
	"gorm.io/gorm"
)

type SubmissionRepository interface {
	GetAll() ([]models.Submission, error)
	AddWorkItem(id uint, wi_id uint) error
	DeleteWorkItem(id uint, wi_id uint) error
}

type submissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}

func (r *submissionRepository) GetAll() ([]models.Submission, error) {
	var subs []models.Submission
	err := r.db.Preload("User").Preload("WorkItems").Find(&subs).Error
	return subs, err
}

func (r *submissionRepository) AddWorkItem(id uint, wi_id uint) error {
	var sItem models.Submission
	var wItem models.WorkItem
	if err := r.db.First(&sItem, id).Error; err != nil {
		return err
	}

	if err := r.db.First(&wItem, wi_id).Error; err != nil {
		return err
	}

	return r.db.Model(&sItem).Association("WorkItems").Append(&wItem)
}

func (r *submissionRepository) DeleteWorkItem(id uint, wi_id uint) error {
	var sItem models.Submission
	var wItem models.WorkItem

	if err := r.db.First(&sItem, id).Error; err != nil {
		return err
	}

	if err := r.db.First(&wItem, wi_id).Error; err != nil {
		return err
	}

	return r.db.Model(&sItem).Association("WorkItems").Delete(&wItem)
}
