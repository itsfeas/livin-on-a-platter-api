package queue_model

import (
	"github.com/google/uuid"
)

type QueuedImage struct {
	UploadId uuid.UUID `json:"id"`
	ImageId  uuid.UUID `json:"image_id"`
}
