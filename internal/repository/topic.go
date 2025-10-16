package repository

import (
	"akademi-business-case/entity"
	"akademi-business-case/model"

	"gorm.io/gorm"
)

type ITopicRepository interface {
	CreateTopic(tx *gorm.DB, topic *entity.Topic) (*entity.Topic, error)
	GetTopic(tx *gorm.DB, param model.GetTopicParam) (*entity.Topic, error)
}

type TopicRepository struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) ITopicRepository {
	return &TopicRepository{db: db}
}

func (r *TopicRepository) CreateTopic(tx *gorm.DB, topic *entity.Topic) (*entity.Topic, error) {
	err := tx.Debug().Create(&topic).Error
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func (r *TopicRepository) GetTopic(tx *gorm.DB, param model.GetTopicParam) (*entity.Topic, error) {
	var topic *entity.Topic
	err := tx.Debug().Where(&param).First(&topic).Error
	if err != nil {
		return nil, err
	}

	return topic, nil
}
