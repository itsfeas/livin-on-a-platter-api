package model

import (
	"time"

	"github.com/google/uuid"
)

type ImageUpload struct {
	ID        uuid.UUID `json:"id"`
	FileType  string    `json:"file"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
