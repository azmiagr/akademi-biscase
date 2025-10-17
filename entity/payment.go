package entity

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	PaymentID       uuid.UUID `json:"payment_id" gorm:"type:varchar(36);primaryKey"`
	UserID          uuid.UUID `json:"user_id"`
	ClassID         uuid.UUID `json:"class_id"`
	Amount          float64   `json:"amount" gorm:"type:decimal(10,2);not null"`
	Status          string    `json:"status" gorm:"type:enum('pending','paid','failed','expired');default:'pending'"`
	SnapURL         string    `json:"snap_url" gorm:"type:varchar(255);default:null"`
	MidtransOrderID string    `json:"midtrans_payment_id" gorm:"type:varchar(255);default:null;index"`
	PaymentMethod   string    `json:"payment_method" gorm:"type:varchar(255);"`
	CreatedAt       time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
