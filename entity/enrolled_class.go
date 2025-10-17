package entity

import (
	"time"

	"github.com/google/uuid"
)

type EnrolledClass struct {
	UserID    uuid.UUID `json:"user_id" gorm:"type:varchar(36);primaryKey"`
	ClassID   uuid.UUID `json:"class_id" gorm:"type:varchar(36);primaryKey"`
	Status    string    `json:"status" gorm:"type:enum('enrolled','completed');default:'enrolled'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
