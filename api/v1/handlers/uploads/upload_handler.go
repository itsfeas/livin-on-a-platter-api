package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"livin-on-a-platter-api/internal/db/storage"
	img_model "livin-on-a-platter-api/internal/model/img"
	msg "livin-on-a-platter-api/internal/model/msg/types"
	queue_model "livin-on-a-platter-api/internal/model/queue"
	"livin-on-a-platter-api/internal/repository"
	http_util "livin-on-a-platter-api/internal/util/error"
	"mime/multipart"
	"net/http"
	"os"
)

const (
	MaxUploadSize = 10_000_000
	MaxMemSize    = 1_000_000
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	bucket := os.Getenv("IMG_STORAGE_BUCKET")
	if r.Method != http.MethodPost {
		http_util.WriteError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MaxUploadSize)
	if err := r.ParseMultipartForm(MaxMemSize); err != nil {
		http_util.WriteError(w, "file over the maximum size allowed", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http_util.WriteError(w, "Unable to retrieve the file from the form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imgId := uuid.New()
	err, success := uploadFile(w, err, file, bucket, imgId)
	if success {
		return
	}

	uploadId := uuid.New()
	img := &img_model.Image{
		ID:       imgId,
		FileType: "png",
		Version:  0,
	}

	err, success = writeImageUploadEntry(w, uploadId, img)
	if success {
		return
	}

	err, success = queueImage(w, uploadId, imgId)
	if success {
		return
	}

	respond(w, uploadId)
}

func uploadFile(w http.ResponseWriter, err error, file multipart.File, bucket string, imgId uuid.UUID) (error, bool) {
	storageClient := storage.GetStorage()
	err = storageClient.StreamFileUpload(&file, bucket, imgId.String())
	if err != nil {
		http_util.WriteError(
			w,
			fmt.Sprintf("Error streaming file upload: %v", err),
			http.StatusInternalServerError,
		)
		return nil, true
	}
	return err, false
}

func writeImageUploadEntry(w http.ResponseWriter, uploadId uuid.UUID, img *img_model.Image) (error, bool) {
	imgRepo := repository.NewImageUploadRepository()
	err := imgRepo.Create(img_model.NewImageUpload(uploadId, img))
	if err != nil {
		http_util.WriteError(
			w,
			fmt.Sprintf("Error while queueing image: %v", err),
			http.StatusInternalServerError,
		)
		return nil, true
	}
	return err, false
}

func queueImage(w http.ResponseWriter, uploadId uuid.UUID, imgId uuid.UUID) (error, bool) {
	queueRepo := repository.NewQueueRepository()
	err := queueRepo.Create(&queue_model.QueuedImage{
		UploadId: uploadId,
		ImageId:  imgId,
	})
	if err != nil {
		// TODO: Cleanup sequence for ImageRepo document creation
		http_util.WriteError(
			w,
			fmt.Sprintf("Error while queueing image: %v", err),
			http.StatusInternalServerError,
		)
		return err, true
	}
	return err, false
}

func respond(w http.ResponseWriter, uploadId uuid.UUID) {
	// Create a generic success response
	resp := msg.DefaultDataMsg()
	resp.Data["id"] = uploadId.String()

	// Convert the success response to JSON
	jsonResp, err := resp.ToJson()
	if err != nil {
		http_util.WriteError(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonResp); err != nil {
		http_util.WriteError(w, "Error writing JSON", http.StatusInternalServerError)
		return
	}
}
