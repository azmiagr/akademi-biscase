package service

import (
	"akademi-business-case/internal/repository"
	"akademi-business-case/pkg/bcrypt"
	"akademi-business-case/pkg/jwt"
)

type Service struct {
}

func NewService(repository *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{}
}
