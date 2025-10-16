package repository

import (
	"akademi-business-case/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IContentRepository interface {
	CreateContent(tx *gorm.DB, content *entity.Content) (*entity.Content, error)
	CountContentByTopicID(tx *gorm.DB, topicID uuid.UUID) (int, error)
}

type ContentRepository struct {
	db *gorm.DB
}

func NewContentRepository(db *gorm.DB) IContentRepository {
	return &ContentRepository{db: db}
}

func (r *ContentRepository) CreateContent(tx *gorm.DB, content *entity.Content) (*entity.Content, error) {
	err := tx.Debug().Create(&content).Error
	if err != nil {
		return nil, err
	}

	return content, nil
}

func (r *ContentRepository) CountContentByTopicID(tx *gorm.DB, topicID uuid.UUID) (int, error) {
	var count int64
	err := tx.Debug().Model(&entity.Content{}).Where("topic_id = ?", topicID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}
