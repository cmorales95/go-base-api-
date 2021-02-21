package model

import (
	"time"

	"github.com/google/uuid"
)

// Base is a model to initialize other models
type Base struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
