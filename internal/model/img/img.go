package img

import (
	"time"

	"github.com/google/uuid"
)

type ImageUpload struct {
	ID        uuid.UUID `json:"id"`
	FileType  string    `json:"file_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
