package entity

import (
	"time"

	"github.com/google/uuid"
)

type Topic struct {
	TopicID   uuid.UUID `json:"topic_id" gorm:"type:varchar(36);primaryKey"`
	ClassID   uuid.UUID `json:"class_id"`
	Name      string    `json:"name" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Contents []Content `gorm:"foreignKey:TopicID"`
}
