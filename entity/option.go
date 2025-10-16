package entity

import "github.com/google/uuid"

type Option struct {
	OptionID   uuid.UUID `json:"option_id" gorm:"type:varchar(36);primaryKey"`
	QuestionID uuid.UUID `json:"question_id"`
	OptionText string    `json:"option_text" gorm:"type:varchar(255);not null"`
	IsCorrect  bool      `json:"is_correct" gorm:"type:boolean;default:false"`
}
