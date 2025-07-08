package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Soft delete field, if not null, the record is considered deleted
	// This field is optional and can be used to implement soft deletes
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"type:timestamp;default:null"`
	IsDeleted bool       `json:"-" gorm:"column:is_deleted;default:false"`
}
