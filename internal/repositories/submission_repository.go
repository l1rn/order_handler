package repositories

import (
	"github.com/l1rn/order-handler/internal/models"
	"gorm.io/gorm"
)

type SubmissionRepository interface {
	GetAll() ([]models.Submission, error)
}

type submissionRepository struct {
	db *gorm.DB
}

func NewSubmissionRepository(db *gorm.DB) SubmissionRepository {
	return &submissionRepository{db: db}
}

func (r *submissionRepository) GetAll() ([]models.Submission, error) {
	var subs []models.Submission
	err := r.db.Find(&subs).Error
	return subs, err
}
