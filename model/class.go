package model

import "github.com/google/uuid"

type GetAllClassesResponse struct {
	ClassID     uuid.UUID `json:"product_id" gorm:"type:varchar(36);primaryKey"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`
	Description string    `json:"description" gorm:"type:text;default:null"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	Discount    float64   `json:"discount" gorm:"type:decimal(10,2);default:0"`
	ImageURL    string    `json:"image_url" gorm:"type:varchar(255);default:null"`
	TotalRating float64   `json:"total_ratings" gorm:"type:decimal(2,1);default:0"`
	TotalReview int       `json:"total_reviews" gorm:"type:int;default:0"`
}

type ClassParam struct {
	ClassID     uuid.UUID `json:"-"`
	ClassTypeID uuid.UUID `json:"-"`
	Name        string    `json:"-"`
}

type GetClassResponse struct {
	ClassID     uuid.UUID `json:"product_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Discount    float64   `json:"discount"`
	ImageURL    string    `json:"image_url"`
	TotalRating float64   `json:"total_ratings"`
	TotalReview int       `json:"total_reviews"`
}

type CreateClassRequest struct {
	Name        string    `json:"name" binding:"required"`
	UserID      uuid.UUID `json:"mentor_id" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Price       float64   `json:"price" binding:"required"`
	Discount    float64   `json:"discount"`
	ImageURL    string    `json:"image_url" binding:"required"`
}

type CreateClassResponse struct {
	ClassID uuid.UUID `json:"class_id"`
}

type GetClassAdminResponse struct {
	ClassName        string         `json:"class_name"`
	ClassDescription string         `json:"class_description"`
	ClassContents    []ClassContent `json:"class_contents"`
}

type ClassContent struct {
	TopicName        string                 `json:"topic_name"`
	ContentResponses []ClassContentResponse `json:"content_responses"`
}

type ClassContentResponse struct {
	ContentTitle string `json:"content_title"`
}
