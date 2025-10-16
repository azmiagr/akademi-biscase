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
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		UserService:    NewUserService(repository.UserRepository, repository.CartRepository, repository.OtpRepository, bcrypt, jwtAuth),
		OtpService:     NewOtpService(repository.OtpRepository, repository.UserRepository),
		ClassService:   NewClassService(repository.ClassRepository, repository.UserRepository),
		ContentService: NewContentService(repository.ContentRepository, repository.TopicRepository, repository.ClassRepository),
	}
}
