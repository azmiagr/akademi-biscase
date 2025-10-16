package entity

import (
	"time"

	"github.com/google/uuid"
)

type ContentType string

const (
	Video ContentType = "video"
	Quiz  ContentType = "quiz"
)

type Content struct {
	ContentID   uuid.UUID   `json:"content_id" gorm:"type:varchar(36);primaryKey"`
	TopicID     uuid.UUID   `json:"topic_id"`
	Title       string      `json:"title" gorm:"type:varchar(255);not null"`
	Type        ContentType `json:"type" gorm:"type:enum('video','quiz','summary','exam');default:'video'"`
	Description string      `json:"description" gorm:"type:text;default:null"`
	ContentURL  string      `json:"content_url" gorm:"type:varchar(255);default:null"`
	Sequencence int         `json:"sequencence" gorm:"type:int;default:0"`
	CreatedAt   time.Time   `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `json:"updated_at" gorm:"autoUpdateTime"`

	Questions []Question `json:"questions,omitempty" gorm:"foreignKey:ContentID"`
}
