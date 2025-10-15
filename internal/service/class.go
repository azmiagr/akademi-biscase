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

type IClassService interface {
	GetAllClasses() ([]*model.GetAllClassesResponse, error)
	GetClassDetail(classID uuid.UUID) (*model.GetClassResponse, error)
	GetClassesByType(classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error)
	GetClassByName(name string, classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error)
	CreateClass(param *model.CreateClassRequest) (*model.CreateClassResponse, error)
}

type ClassService struct {
	db              *gorm.DB
	ClassRepository repository.IClassRepository
	UserRepository  repository.IUserRepository
}

func NewClassService(classRepository repository.IClassRepository, userRepository repository.IUserRepository) IClassService {
	return &ClassService{
		db:              mariadb.Connection,
		ClassRepository: classRepository,
		UserRepository:  userRepository,
	}
}

func (s *ClassService) GetAllClasses() ([]*model.GetAllClassesResponse, error) {
	var result []*model.GetAllClassesResponse

	tx := s.db.Begin()
	defer tx.Rollback()

	classes, err := s.ClassRepository.GetAllClasses(tx)
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

func (s *ClassService) GetClassDetail(classID uuid.UUID) (*model.GetClassResponse, error) {
	var result *model.GetClassResponse

	tx := s.db.Begin()
	defer tx.Rollback()

	class, err := s.ClassRepository.GetClass(tx, model.ClassParam{
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

func (s *ClassService) GetClassesByType(classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error) {
	var result []*model.GetAllClassesResponse

	tx := s.db.Begin()
	defer tx.Rollback()

	classes, err := s.ClassRepository.GetClassesByType(tx, model.ClassParam{
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

func (s *ClassService) GetClassByName(name string, classTypeID uuid.UUID) ([]*model.GetAllClassesResponse, error) {
	var result []*model.GetAllClassesResponse

	tx := s.db.Begin()
	defer tx.Rollback()

	classes, err := s.ClassRepository.GetClassByName(tx, model.ClassParam{
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

func (s *ClassService) CreateClass(param *model.CreateClassRequest) (*model.CreateClassResponse, error) {
	tx := s.db.Begin()
	defer tx.Rollback()

	classID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	user, err := s.UserRepository.GetUser(tx, model.UserParam{
		UserID: param.UserID,
	})
	if err != nil {
		return nil, errors.New("user not found")
	}

	if user.RoleID != 3 {
		return nil, errors.New("user is not mentor")
	}

	class := &entity.Class{
		ClassID:     classID,
		UserID:      user.UserID,
		Name:        param.Name,
		Description: param.Description,
		Price:       param.Price,
		Discount:    param.Discount,
		ImageURL:    param.ImageURL,
	}

	_, err = s.ClassRepository.CreateClass(tx, class)
	if err != nil {
		return nil, err
	}

	result := &model.CreateClassResponse{
		ClassID: classID,
	}

	err = tx.Commit().Error
	if err != nil {
		return nil, err
	}

	return result, nil
}
