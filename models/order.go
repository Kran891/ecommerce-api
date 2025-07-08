package models

import "github.com/google/uuid"

type Order struct {
	UserID  uuid.UUID `json:"userID"`
	
}
