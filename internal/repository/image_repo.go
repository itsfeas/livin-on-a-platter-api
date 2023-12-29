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

const imgDocPrefix string = "images/"
const imgLogPrefix string = "img-repo | "

func NewImageRepository() *ImageRepository {
	d := firebase.GetDB()
	return &ImageRepository{
		FireDB: d,
	}
}

func (i *ImageRepository) Create(image *img_model.Image) error {
	ref := i.NewRef(imgDocPrefix + image.ID.String())
	if err := ref.Set(context.Background(), image); err != nil {
		return fmt.Errorf("%s couldn't CREATE image entry %s: %v", imgLogPrefix, image.ID.String(), err)
	}
	return nil
}

func (i *ImageRepository) Delete(id string) error {
	ref := i.NewRef(imgDocPrefix + id)
	if err := ref.Delete(context.Background()); err != nil {
		return fmt.Errorf("%s couldn't DELETE image entry %s: %v", imgLogPrefix, id, err)
	}
	return nil
}

func (i *ImageRepository) Update(upload *img_model.ImageUpload) error {
	id := upload.UploadId.String()
	ref := i.NewRef(imgDocPrefix + id)
	if err := ref.Set(context.Background(), upload); err != nil {
		return fmt.Errorf("%s couldn't UPDATE image entry %s: %v", imgLogPrefix, id, err)
	}
	return nil
}

func (i *ImageRepository) FindById(id string) (*img_model.ImageUpload, error) {
	upload := &img_model.ImageUpload{}
	ref := i.NewRef(imgDocPrefix + id)
	err := ref.Get(context.Background(), upload)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID upload %s: %v", imgLogPrefix, id, err)
	}
	if upload.UploadId.String() == "" {
		return nil, nil
	}
	return upload, nil
}

func (i *ImageRepository) Find(upload *img_model.ImageUpload) (*img_model.ImageUpload, error) {
	id := upload.UploadId.String()
	ref := i.NewRef(imgDocPrefix + id)
	err := ref.Get(context.Background(), upload)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID image entry %s: %v", imgLogPrefix, id, err)
	}
	if upload.UploadId.String() == "" {
		return nil, nil
	}
	return upload, nil
}
