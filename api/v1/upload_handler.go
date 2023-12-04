package api

import (
	"encoding/json"
	"fmt"
	"livin-on-a-platter-api/internal/model"
	"livin-on-a-platter-api/internal/repository"
	"livin-on-a-platter-api/internal/storage"
	"net/http"
	"time"

	"github.com/google/uuid"
)

const MAX_UPLOAD_SIZE = 10_000_000
const MAX_MEM_SIZE = 1_000_000

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_MEM_SIZE); err != nil {
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

	storage_client := storage.GetStorage()
	err = storage_client.StreamFileUpload(&file, "loap-img-storage", uuid.String())
	if err != nil {
		http.Error(w, fmt.Sprintf("Error streaming file upload: %v", err), http.StatusInternalServerError)
		return
	}

	repo := repository.NewImageRepository()
	repo.Create(&model.ImageUpload{
		ID:        uuid,
		FileType:  "png",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	// Create a generic success response
	resp := map[string]interface{}{
		"status": "ok",
		"id":     uuid.String(),
	}

	// Convert the success response to JSON
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
