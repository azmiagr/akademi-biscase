package entity

import "github.com/google/uuid"

type Class struct {
	ClassID     uuid.UUID `json:"product_id" gorm:"type:varchar(36);primaryKey"`
	ClassTypeID uuid.UUID `json:"class_type_id"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;default:null"`
	Price       float64   `json:"price" gorm:"type:decimal(8,2);not null"`
	ImageURL    string    `json:"image_url" gorm:"type:varchar(255);default:null"`
	Benefit     string    `json:"benefit" gorm:"type:text;default:null"`

	CartItems []CartItem `gorm:"foreignKey:ClassID"`
}
