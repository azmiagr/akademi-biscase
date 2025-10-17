package service

import (
	"akademi-business-case/entity"
	"akademi-business-case/internal/repository"
	"akademi-business-case/model"
	"akademi-business-case/pkg/database/mariadb"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IContentService interface {
	CreateContent(param *model.CreateContentRequest, classID uuid.UUID) (*model.GetClassAdminResponse, error)
}

type ContentService struct {
	db                *gorm.DB
	ContentRepository repository.IContentRepository
	TopicRepository   repository.ITopicRepository
	ClassRepository   repository.IClassRepository
}

func NewContentService(contentRepository repository.IContentRepository, topicRepository repository.ITopicRepository, classRepository repository.IClassRepository) IContentService {
	return &ContentService{
		db:                mariadb.Connection,
		ContentRepository: contentRepository,
		TopicRepository:   topicRepository,
		ClassRepository:   classRepository,
	}
}

func (s *ContentService) CreateContent(param *model.CreateContentRequest, classID uuid.UUID) (*model.GetClassAdminResponse, error) {
	var (
		result        *model.GetClassAdminResponse
		existingTopic *entity.Topic
	)

	tx := s.db.Begin()
	defer tx.Rollback()

	class, err := s.ClassRepository.GetClass(tx, model.ClassParam{
		ClassID: classID,
	})
	if err != nil {
		return nil, errors.New("class not found")
	}

	existingTopic, err = s.TopicRepository.FindByNameAndClassID(tx, param.TopicName, class.ClassID)
	if err != nil {
		topicID, err := uuid.NewUUID()
		if err != nil {
			return nil, err
		}

		newTopic := &entity.Topic{
			TopicID: topicID,
			ClassID: class.ClassID,
			Name:    param.TopicName,
		}
		createdTopic, err := s.TopicRepository.CreateTopic(tx, newTopic)
		if err != nil {
			return nil, err
		}
		existingTopic = createdTopic
	}

	contentID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	count, err := s.ContentRepository.CountContentByTopicID(tx, existingTopic.TopicID)
	if err != nil {
		return nil, err
	}

	content := &entity.Content{
		ContentID:   contentID,
		TopicID:     existingTopic.TopicID,
		Title:       param.Title,
		Type:        entity.ContentType(param.ContentType),
		Description: param.Description,
		ContentURL:  param.URL,
		Sequencence: count + 1,
	}

	_, err = s.ContentRepository.CreateContent(tx, content)
	if err != nil {
		return nil, err
	}

	class, err = s.ClassRepository.GetClass(tx, model.ClassParam{
		ClassID: classID,
	})
	if err != nil {
		return nil, err
	}

	classContents := []model.ClassContent{}
	for _, topic := range class.Topics {
		classContent := model.ClassContent{
			TopicName:        topic.Name,
			ContentResponses: []model.ClassContentResponse{},
		}
		for _, content := range topic.Contents {
			contentResponse := model.ClassContentResponse{
				ContentTitle: content.Title,
			}
			classContent.ContentResponses = append(classContent.ContentResponses, contentResponse)
		}
		classContents = append(classContents, classContent)
	}

	result = &model.GetClassAdminResponse{
		ClassName:        class.Name,
		ClassDescription: class.Description,
		ClassContents:    classContents,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return result, nil

}
