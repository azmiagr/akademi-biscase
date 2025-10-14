package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository  IUserRepository
	OtpRepository   IOtpRepository
	CartRepository  ICartRepository
	ClassRepository IClassRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:  NewUserRepository(db),
		OtpRepository:   NewOtpRepository(db),
		CartRepository:  NewCartRepository(db),
		ClassRepository: NewClassRepository(db),
	}
}
