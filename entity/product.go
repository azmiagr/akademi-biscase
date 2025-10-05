package entity

import "github.com/google/uuid"

type Product struct {
	ProductID   uuid.UUID `json:"product_id" gorm:"type:varchar(36);primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;default:null"`
	Price       float64   `json:"price" gorm:"type:decimal(8,2);not null"`
	ImageURL    string    `json:"image_url" gorm:"type:varchar(255);default:null"`

	CartItems []CartItem `gorm:"foreignKey:ProductID"`
}
