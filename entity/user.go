package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `json:"user_id" gorm:"type:varchar(36);primaryKey"`
	RoleID    int       `json:"role_id"`
	FirstName string    `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  string    `json:"last_name" gorm:"type:varchar(100);not null"`
	Email     string    `json:"email" gorm:"type:varchar(50);uniqueIndex;not null"`
	Password  string    `json:"password" gorm:"type:varchar(80);not null"`
	Status    string    `json:"status" gorm:"type:enum('active','inactive');default:'inactive'"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`

	Cart     Cart      `gorm:"foreignKey:UserID"`
	OtpCodes []OtpCode `gorm:"foreignKey:UserID"`
	Reviews  []Review  `gorm:"foreignKey:UserID"`
	Classes  []Class   `gorm:"foreignKey:UserID"`
}
