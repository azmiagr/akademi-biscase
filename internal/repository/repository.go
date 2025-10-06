package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	OtpRepository  IOtpRepository
	CartRepository ICartRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
		OtpRepository:  NewOtpRepository(db),
		CartRepository: NewCartRepository(db),
	}
}
