package entity

import (
	"time"

	"github.com/google/uuid"
)

type CartItem struct {
	CartItemID uuid.UUID `json:"cartItemID" gorm:"type:varchar(36);primaryKey"`
	CartID     uuid.UUID `json:"cartID" gorm:"type:varchar(36)"`
	ProductID  uuid.UUID `json:"productID" gorm:"type:varchar(36)"`
	CreatedAt  time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
