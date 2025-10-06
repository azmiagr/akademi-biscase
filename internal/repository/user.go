package repository

import (
	"akademi-business-case/entity"
	"akademi-business-case/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(tx *gorm.DB, user *entity.User) (*entity.User, error)
	GetUser(tx *gorm.DB, param model.UserParam) (*entity.User, error)
	UpdateUser(tx *gorm.DB, user *entity.User) (*entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	err := tx.Debug().Create(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) GetUser(tx *gorm.DB, param model.UserParam) (*entity.User, error) {
	user := entity.User{}
	err := tx.Debug().Preload("Cart").Where(&param).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(tx *gorm.DB, user *entity.User) (*entity.User, error) {
	err := tx.Debug().Where("user_id = ?", user.UserID).Updates(&user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
