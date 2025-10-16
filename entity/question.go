package entity

import (
	"github.com/google/uuid"
)

type Question struct {
	QuestionID   uuid.UUID `json:"question_id" gorm:"type:varchar(36);primaryKey"`
	ContentID    uuid.UUID `json:"content_id"`
	QuestionText string    `json:"question_text" gorm:"type:text;not null"`
	Sequencence  int       `json:"sequencence" gorm:"type:int;default:0"`

	Options []Option `json:"options" gorm:"foreignKey:QuestionID"`
}
