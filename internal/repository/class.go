package repository

import (
	"akademi-business-case/entity"
	"akademi-business-case/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IClassRepository interface {
	GetAllClasses(tx *gorm.DB) ([]*entity.Class, error)
	GetClass(tx *gorm.DB, param model.ClassParam) (*entity.Class, error)
	GetClassesByType(tx *gorm.DB, param model.ClassParam) ([]*entity.Class, error)
	GetClassByName(tx *gorm.DB, param model.ClassParam) ([]*entity.Class, error)
	CreateClass(tx *gorm.DB, class *entity.Class) (*entity.Class, error)
}

type ClassRepository struct {
	db *gorm.DB
}

func NewClassRepository(db *gorm.DB) IClassRepository {
	return &ClassRepository{db: db}
}

func (r *ClassRepository) GetAllClasses(tx *gorm.DB) ([]*entity.Class, error) {
	classes := []*entity.Class{}
	err := tx.Debug().Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (r *ClassRepository) GetClass(tx *gorm.DB, param model.ClassParam) (*entity.Class, error) {
	var class *entity.Class
	err := tx.Debug().Preload("Topics.Contents").Where(&param).First(&class).Error
	if err != nil {
		return nil, err
	}

	return class, nil
}

func (r *ClassRepository) GetClassesByType(tx *gorm.DB, param model.ClassParam) ([]*entity.Class, error) {
	var classes []*entity.Class
	err := tx.Debug().Where(&param).Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (r *ClassRepository) GetClassByName(tx *gorm.DB, param model.ClassParam) ([]*entity.Class, error) {
	var classes []*entity.Class
	query := tx.Debug()

	if param.Name != "" {
		query = query.Where("name LIKE ?", "%"+param.Name+"%")
	}

	if param.ClassTypeID != uuid.Nil {
		query = query.Where("class_type_id = ?", param.ClassTypeID)
	}

	err := query.Find(&classes).Error
	if err != nil {
		return nil, err
	}

	return classes, nil
}

func (r *ClassRepository) CreateClass(tx *gorm.DB, class *entity.Class) (*entity.Class, error) {
	err := tx.Debug().Create(&class).Error
	if err != nil {
		return nil, err
	}

	return class, nil
}
