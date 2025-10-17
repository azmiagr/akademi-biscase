package model

import (
	"time"

	"github.com/google/uuid"
)

type AddReviewRequest struct {
	Rating  int    `json:"rating" binding:"required,min=1,max=5"`
	Comment string `json:"comment" binding:"required"`
}

type AddReviewResponse struct {
	ReviewID  uuid.UUID `json:"review_id"`
	FullName  string    `json:"full_name"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
