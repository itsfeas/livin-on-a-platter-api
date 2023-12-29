package img_model

import (
	"time"

	"github.com/google/uuid"
)

type ImageUpload struct {
	UploadId        uuid.UUID    `json:"id"`
	Image           *uuid.UUID   `json:"image"`
	GeneratedImages []*uuid.UUID `json:"generated_images"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	Version         uint16       `json:"version"`
}

type Image struct {
	ID       uuid.UUID `json:"id"`
	FileType string    `json:"file_type"`
	Version  uint16    `json:"version"`
}

func NewImageUpload(uploadId uuid.UUID, img *Image) *ImageUpload {
	return &ImageUpload{
		UploadId:        uploadId,
		Image:           &img.ID,
		GeneratedImages: make([]*uuid.UUID, 0),
		CreatedAt:       time.Now().UTC(),
		UpdatedAt:       time.Now().UTC(),
		Version:         0,
	}
}
