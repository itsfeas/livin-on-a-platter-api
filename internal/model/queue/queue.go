package queue_model

import (
	"github.com/google/uuid"
)

type QueuedImage struct {
	UploadId uuid.UUID `json:"upload_id"`
	ImageId  uuid.UUID `json:"image_id"`
}
