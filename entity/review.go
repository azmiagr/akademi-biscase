package entity

import (
	"time"

	"github.com/google/uuid"
)

type Review struct {
	ReviewID  uuid.UUID `json:"review_id" gorm:"type:varchar(36);primaryKey"`
	UserID    uuid.UUID `json:"user_id" gorm:"type:varchar(36);not null"`
	ClassID   uuid.UUID `json:"class_id" gorm:"type:varchar(36);not null"`
	Rating    int       `json:"rating" gorm:"type:int;not null"`
	Comment   string    `json:"comment" gorm:"type:text;default:null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
