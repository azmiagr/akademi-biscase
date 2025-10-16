package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository    IUserRepository
	OtpRepository     IOtpRepository
	CartRepository    ICartRepository
	ClassRepository   IClassRepository
	TopicRepository   ITopicRepository
	ContentRepository IContentRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepository:    NewUserRepository(db),
		OtpRepository:     NewOtpRepository(db),
		CartRepository:    NewCartRepository(db),
		ClassRepository:   NewClassRepository(db),
		TopicRepository:   NewTopicRepository(db),
		ContentRepository: NewContentRepository(db),
	}
}
