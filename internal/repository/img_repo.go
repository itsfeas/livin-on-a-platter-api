package repository

import (
	"context"
	"fmt"
	"livin-on-a-platter-api/internal/db/firebase"
	"livin-on-a-platter-api/internal/model/img"
)

type ImageRepository struct {
	*firebase.FireDB
}

const DOC_PREFIX string = "img_upload/"
const LOG_PREFIX string = "img-repo | "

func NewImageRepository() *ImageRepository {
	d := firebase.GetDB()
	return &ImageRepository{
		FireDB: d,
	}
}

func (i *ImageRepository) Create(upload *img.ImageUpload) error {
	ref := i.NewRef(DOC_PREFIX + upload.ID.String())
	if err := ref.Set(context.Background(), upload); err != nil {
		return fmt.Errorf("%s couldn't CREATE upload %s: %v", LOG_PREFIX, upload.ID.String(), err)
	}
	return nil
}

func (i *ImageRepository) Delete(id string) error {
	ref := i.NewRef(DOC_PREFIX + id)
	if err := ref.Delete(context.Background()); err != nil {
		return fmt.Errorf("%s couldn't DELETE upload %s: %v", LOG_PREFIX, id, err)
	}
	return nil
}

func (i *ImageRepository) Update(upload *img.ImageUpload) error {
	id := upload.ID.String()
	ref := i.NewRef(DOC_PREFIX + id)
	if err := ref.Set(context.Background(), upload); err != nil {
		return fmt.Errorf("%s couldn't UPDATE upload %s: %v", LOG_PREFIX, id, err)
	}
	return nil
}

func (i *ImageRepository) FindById(id string) (*img.ImageUpload, error) {
	upload := &img.ImageUpload{}
	ref := i.NewRef(DOC_PREFIX + id)
	err := ref.Get(context.Background(), upload)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID upload %s: %v", LOG_PREFIX, id, err)
	}
	if upload.ID.String() == "" {
		return nil, nil
	}
	return upload, nil
}

func (i *ImageRepository) Find(upload *img.ImageUpload) (*img.ImageUpload, error) {
	id := upload.ID.String()
	ref := i.NewRef(DOC_PREFIX + id)
	err := ref.Get(context.Background(), upload)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID upload %s: %v", LOG_PREFIX, id, err)
	}
	if upload.ID.String() == "" {
		return nil, nil
	}
	return upload, nil
}
