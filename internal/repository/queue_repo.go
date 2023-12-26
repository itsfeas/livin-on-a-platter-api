package repository

import (
	"context"
	"fmt"
	"livin-on-a-platter-api/internal/db/firebase"
	queue_model "livin-on-a-platter-api/internal/model/queue"
)

type QueueRepository struct {
	*firebase.FireDB
}

const queueDocPrefix string = "queue/"
const queueLogPrefix string = "queue-repo | "

func NewQueueRepository() *QueueRepository {
	d := firebase.GetDB()
	return &QueueRepository{
		FireDB: d,
	}
}

func (i *QueueRepository) Create(queueItem *queue_model.QueuedImage) error {
	ref := i.NewRef(queueDocPrefix + queueItem.UploadId.String())
	if err := ref.Set(context.Background(), queueItem); err != nil {
		return fmt.Errorf("%s couldn't CREATE queue_item %s: %v", queueLogPrefix, queueItem.UploadId.String(), err)
	}
	return nil
}

func (i *QueueRepository) Delete(id string) error {
	ref := i.NewRef(queueDocPrefix + id)
	if err := ref.Delete(context.Background()); err != nil {
		return fmt.Errorf("%s couldn't DELETE queue_item %s: %v", queueLogPrefix, id, err)
	}
	return nil
}

func (i *QueueRepository) Update(queueItem *queue_model.QueuedImage) error {
	id := queueItem.UploadId.String()
	ref := i.NewRef(queueDocPrefix + id)
	if err := ref.Set(context.Background(), queueItem); err != nil {
		return fmt.Errorf("%s couldn't UPDATE queue_item %s: %v", queueLogPrefix, id, err)
	}
	return nil
}

func (i *QueueRepository) FindById(id string) (*queue_model.QueuedImage, error) {
	upload := &queue_model.QueuedImage{}
	ref := i.NewRef(queueDocPrefix + id)
	err := ref.Get(context.Background(), upload)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID queue_item %s: %v", queueLogPrefix, id, err)
	}
	if upload.UploadId.String() == "" {
		return nil, nil
	}
	return upload, nil
}

func (i *QueueRepository) Find(queueItem *queue_model.QueuedImage) (*queue_model.QueuedImage, error) {
	id := queueItem.UploadId.String()
	ref := i.NewRef(queueDocPrefix + id)
	err := ref.Get(context.Background(), queueItem)
	if err != nil {
		return nil, fmt.Errorf("%s couldn't FIND_BY_ID upload %s: %v", queueLogPrefix, id, err)
	}
	if queueItem.UploadId.String() == "" {
		return nil, nil
	}
	return queueItem, nil
}
