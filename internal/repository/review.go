package repository

import (
	"akademi-business-case/entity"

	"gorm.io/gorm"
)

type IReviewRepository interface {
	CreateReview(tx *gorm.DB, review *entity.Review) (*entity.Review, error)
}

type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) IReviewRepository {
	return &ReviewRepository{db: db}
}

func (r *ReviewRepository) CreateReview(tx *gorm.DB, review *entity.Review) (*entity.Review, error) {
	err := tx.Debug().Create(&review).Error
	if err != nil {
		return nil, err
	}

	return review, nil
}
