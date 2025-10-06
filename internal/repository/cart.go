package repository

import (
	"akademi-business-case/entity"
	"akademi-business-case/model"

	"gorm.io/gorm"
)

type ICartRepository interface {
	CreateCart(tx *gorm.DB, cart *entity.Cart) (*entity.Cart, error)
	GetCart(tx *gorm.DB, param model.CartParam) (*entity.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) ICartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) CreateCart(tx *gorm.DB, cart *entity.Cart) (*entity.Cart, error) {
	err := tx.Debug().Create(&cart).Error
	if err != nil {
		return nil, err
	}

	return cart, nil
}

func (r *CartRepository) GetCart(tx *gorm.DB, param model.CartParam) (*entity.Cart, error) {
	cart := entity.Cart{}
	err := tx.Debug().Preload("CartItems").Where(&param).First(&cart).Error
	if err != nil {
		return nil, err
	}

	return &cart, nil
}
