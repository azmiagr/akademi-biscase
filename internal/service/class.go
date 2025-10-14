package service

import (
	"akademi-business-case/internal/repository"
	"akademi-business-case/model"
	"akademi-business-case/pkg/database/mariadb"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IClassService interface {
	GetAllClasses() ([]*model.GetAllClassesResponse, error)
	GetClassDetail(classID uuid.UUID) (*model.GetClassResponse, error)
	GetClassesByType(classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error)
	GetClassByName(name string, classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error)
}

type ClassService struct {
	db              *gorm.DB
	ClassRepository repository.IClassRepository
}

func NewClassService(classRepository repository.IClassRepository) IClassService {
	return &ClassService{
		db:              mariadb.Connection,
		ClassRepository: classRepository,
	}
}

func (c *ClassService) GetAllClasses() ([]*model.GetAllClassesResponse, error) {
	var result []*model.GetAllClassesResponse

	tx := c.db.Begin()
	defer tx.Rollback()

	classes, err := c.ClassRepository.GetAllClasses(tx)
	if err != nil {
		return nil, err
	}

	for _, class := range classes {
		result = append(result, &model.GetAllClassesResponse{
			ClassID:     class.ClassID,
			Name:        class.Name,
			Description: class.Description,
			Price:       class.Price,
			Discount:    class.Discount,
			ImageURL:    class.ImageURL,
			TotalRating: class.TotalRating,
			TotalReview: class.TotalReview,
		})
	}

	return result, nil
}

func (c *ClassService) GetClassDetail(classID uuid.UUID) (*model.GetClassResponse, error) {
	var result *model.GetClassResponse

	tx := c.db.Begin()
	defer tx.Rollback()

	class, err := c.ClassRepository.GetClass(tx, model.ClassParam{
		ClassID: classID,
	})
	if err != nil {
		return nil, err
	}

	result = &model.GetClassResponse{
		ClassID:     class.ClassID,
		Name:        class.Name,
		Description: class.Description,
		Price:       class.Price,
		Discount:    class.Discount,
		ImageURL:    class.ImageURL,
		TotalRating: class.TotalRating,
		TotalReview: class.TotalReview,
	}

	return result, nil
}

func (c *ClassService) GetClassesByType(classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error) {
	var result []*model.GetAllClassesResponse

	tx := c.db.Begin()
	defer tx.Rollback()

	classes, err := c.ClassRepository.GetClassesByType(tx, model.ClassParam{
		ClassTypeID: classTypeID,
	})
	if err != nil {
		return nil, err
	}

	for _, class := range classes {
		result = append(result, &model.GetAllClassesResponse{
			ClassID:     class.ClassID,
			Name:        class.Name,
			Description: class.Description,
			Price:       class.Price,
			Discount:    class.Discount,
			ImageURL:    class.ImageURL,
			TotalRating: class.TotalRating,
			TotalReview: class.TotalReview,
		})
	}

	return result, nil
}

func (c *ClassService) GetClassByName(name string, classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error) {
	var result []*model.GetAllClassesResponse

	tx := c.db.Begin()
	defer tx.Rollback()

	classes, err := c.ClassRepository.GetClassByName(tx, model.ClassParam{
		Name:        name,
		ClassTypeID: classTypeID,
	})
	if err != nil {
		return nil, err
	}

	for _, class := range classes {
		result = append(result, &model.GetAllClassesResponse{
			ClassID:     class.ClassID,
			Name:        class.Name,
			Description: class.Description,
			Price:       class.Price,
			Discount:    class.Discount,
			ImageURL:    class.ImageURL,
			TotalRating: class.TotalRating,
			TotalReview: class.TotalReview,
		})
	}

	return result, nil
}
