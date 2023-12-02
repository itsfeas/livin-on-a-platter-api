package repository

import (
	"livin-on-a-platter-api/internal/db/firebase"
)

type ImageRepository struct {
	*firebase.FireDB
}
