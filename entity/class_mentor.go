package entity

import "github.com/google/uuid"

type ClassMentor struct {
	ClassID uuid.UUID `json:"class_id" gorm:"type:varchar(36);primaryKey"`
	UserID  uuid.UUID `json:"user_id" gorm:"type:varchar(36);primaryKey"`
}
