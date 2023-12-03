package api

import (
	"fmt"
	"livin-on-a-platter-api/internal/model"
	"livin-on-a-platter-api/internal/repository"
	"livin-on-a-platter-api/internal/storage"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const MAX_UPLOAD_SIZE = 10_000_000

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "file over the maximum size allowed", http.StatusBadRequest)
		return
	}
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to retrieve the file from the form", http.StatusBadRequest)
		return
	}
	defer file.Close()

	uuid := uuid.New()
	repo := repository.NewImageRepository()
	repo.Create(&model.ImageUpload{
		ID:        uuid,
		FileName:  "file",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	storage_client := storage.GetStorage()
	err = storage_client.StreamFileUpload(&file, "loap-img-storage", uuid.String())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error streaming file upload: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
