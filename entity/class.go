package entity

import "github.com/google/uuid"

type Class struct {
	ClassID     uuid.UUID `json:"product_id" gorm:"type:varchar(36);primaryKey"`
	ClassTypeID uuid.UUID `json:"class_type_id"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;default:null"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Discount    float64   `json:"discount" gorm:"type:decimal(10,2);default:0"`
	ImageURL    string    `json:"image_url" gorm:"type:varchar(255);default:null"`
	Benefit     string    `json:"benefit" gorm:"type:text;default:null"`
	TotalRating float64   `json:"total_ratings" gorm:"type:decimal(2,1);default:0"`
	TotalReview int       `json:"total_reviews" gorm:"type:int;default:0"`

	ClassMentors    []ClassMentor   `gorm:"foreignKey:ClassID"`
	Reviews         []Review        `gorm:"foreignKey:ClassID"`
	CartItems       []CartItem      `gorm:"foreignKey:ClassID"`
	Topics          []Topic         `gorm:"foreignKey:ClassID"`
	EnrolledClasses []EnrolledClass `gorm:"foreignKey:ClassID"`
	Payments        []Payment       `gorm:"foreignKey:ClassID"`
}
