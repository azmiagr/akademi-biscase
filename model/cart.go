package model

import "github.com/google/uuid"

type CartParam struct {
	CartID uuid.UUID `json:"-"`
	UserID uuid.UUID `json:"-"`
}
