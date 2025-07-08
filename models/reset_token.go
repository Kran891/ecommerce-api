// models/reset_token.go
package models

import (
	"time"

	"github.com/google/uuid"
)

type ResetToken struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID
	Token     string `gorm:"unique;not null"`
	ExpiresAt time.Time
}
