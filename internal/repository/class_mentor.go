package repository

import (
	"akademi-business-case/entity"

	"gorm.io/gorm"
)

type IClassMentorRepository interface {
	CreateClassMentor(tx *gorm.DB, classMentor *entity.ClassMentor) (*entity.ClassMentor, error)
}

type ClassMentorRepository struct {
	db *gorm.DB
}

func NewClassMentorRepository(db *gorm.DB) IClassMentorRepository {
	return &ClassMentorRepository{db: db}
}

func (r *ClassMentorRepository) CreateClassMentor(tx *gorm.DB, classMentor *entity.ClassMentor) (*entity.ClassMentor, error) {
	err := tx.Debug().Create(&classMentor).Error
	if err != nil {
		return nil, err
	}

	return classMentor, nil
}
