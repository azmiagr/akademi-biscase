package service

import (
	"akademi-business-case/entity"
	"akademi-business-case/internal/repository"
	"akademi-business-case/model"
	"akademi-business-case/pkg/database/mariadb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IReviewService interface {
	AddReview(param *model.AddReviewRequest, userID uuid.UUID, classID uuid.UUID) (*model.AddReviewResponse, error)
}

type ReviewService struct {
	db               *gorm.DB
	ReviewRepository repository.IReviewRepository
	UserRepository   repository.IUserRepository
	ClassRepository  repository.IClassRepository
}

func NewReviewService(reviewRepository repository.IReviewRepository, userRepository repository.IUserRepository, classRepository repository.IClassRepository) IReviewService {
	return &ReviewService{
		db:               mariadb.Connection,
		ReviewRepository: reviewRepository,
		UserRepository:   userRepository,
		ClassRepository:  classRepository,
	}
}

func (s *ReviewService) AddReview(param *model.AddReviewRequest, userID uuid.UUID, classID uuid.UUID) (*model.AddReviewResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	user, err := s.UserRepository.GetUser(tx, model.UserParam{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}

	class, err := s.ClassRepository.GetClass(tx, model.ClassParam{
		ClassID: classID,
	})
	if err != nil {
		return nil, err
	}

	class.TotalReview++

	reviewID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	review := &entity.Review{
		ReviewID: reviewID,
		UserID:   user.UserID,
		ClassID:  class.ClassID,
		Rating:   param.Rating,
		Comment:  param.Comment,
	}

	_, err = s.ReviewRepository.CreateReview(tx, review)
	if err != nil {
		return nil, err
	}

	result := &model.AddReviewResponse{
		ReviewID:  reviewID,
		FullName:  user.FirstName + " " + user.LastName,
		Rating:    param.Rating,
		Comment:   param.Comment,
		CreatedAt: review.CreatedAt,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
