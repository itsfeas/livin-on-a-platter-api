package repository

import (
	"context"
	"fmt"
	"livin-on-a-platter-api/internal/db/firebase"
	"livin-on-a-platter-api/internal/model/img"
)

type ImageUploadRepository struct {
	*firebase.FireDB
}

const docPrefix string = "img_upload/"
const logPrefix string = "img-repo | "

func NewImageUploadRepository() *ImageUploadRepository {
	d := firebase.GetDB()
	return &ImageUploadRepository{
		FireDB: d,
	}
}

func (i *ImageUploadRepository) Create(upload *img.ImageUpload) error {
	ref := i.NewRef(docPrefix + upload.UploadId.String())
	if err := ref.Set(context.Background(), upload); err != nil {
		return fmt.Errorf("%s couldn't CREATE upload %s: %v", logPrefix, upload.UploadId.String(), err)
	}
	return nil
}

func (i *ImageUploadRepository) Delete(id string) error {
	ref := i.NewRef(docPrefix + id)
	if err := ref.Delete(context.Background()); err != nil {
		return fmt.Errorf("%s couldn't DELETE upload %s: %v", logPrefix, id, err)
	}
	return nil
}

func (i *ImageUploadRepository) Update(upload *img.ImageUpload) error {
	id := upload.UploadId.String()
	ref := i.NewRef(docPrefix + id)
	if err := ref.Set(context.Background(), upload); err != nil {
		return fmt.Errorf("%s couldn't UPDATE upload %s: %v", logPrefix, id, err)
	}
	return nil
}

func (i *ImageUploadRepository) FindById(id string) (*img.ImageUpload, error) {
	upload := &img.ImageUpload{}
	ref := i.NewRef(docPrefix + id)
	err := ref.Get(context.Background(), upload)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID upload %s: %v", logPrefix, id, err)
	}
	if upload.UploadId.String() == "" {
		return nil, nil
	}
	return upload, nil
}

func (i *ImageUploadRepository) Find(upload *img.ImageUpload) (*img.ImageUpload, error) {
	id := upload.UploadId.String()
	ref := i.NewRef(docPrefix + id)
	err := ref.Get(context.Background(), upload)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID upload %s: %v", logPrefix, id, err)
	}
	if upload.UploadId.String() == "" {
		return nil, nil
	}
	return upload, nil
}