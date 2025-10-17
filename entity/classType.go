package entity

import "github.com/google/uuid"

type ClassType struct {
	ClassTypeID uuid.UUID `json:"class_type_id" gorm:"type:varchar(36);primaryKey;"`
	Name        string    `json:"name" gorm:"type:varchar(100);not null"`

	Classes []Class `gorm:"foreignKey:ClassTypeID"`
}
