package entity

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	CartID    uuid.UUID `json:"cart_id" gorm:"type:varchar(36);primaryKey"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	CartItems []CartItem `gorm:"foreignKey:CartID"`
}
