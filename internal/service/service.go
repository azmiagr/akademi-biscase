package service

import (
	"akademi-business-case/internal/repository"
	"akademi-business-case/pkg/bcrypt"
	"akademi-business-case/pkg/jwt"
)

type Service struct {
	UserService    IUserService
	OtpService     IOtpService
	ClassService   IClassService
	ContentService IContentService
	ReviewService  IReviewService
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		UserService:    NewUserService(repository.UserRepository, repository.CartRepository, repository.OtpRepository, bcrypt, jwtAuth),
		OtpService:     NewOtpService(repository.OtpRepository, repository.UserRepository),
		ClassService:   NewClassService(repository.ClassRepository, repository.UserRepository, repository.ClassMentorRepository),
		ContentService: NewContentService(repository.ContentRepository, repository.TopicRepository, repository.ClassRepository),
		ReviewService:  NewReviewService(repository.ReviewRepository, repository.UserRepository, repository.ClassRepository),
	}
}
