package main

import (
	"fmt"
	"livin-on-a-platter-api/api/v1"
	"livin-on-a-platter-api/internal/db/firebase"
	"livin-on-a-platter-api/internal/model/img"
	"livin-on-a-platter-api/internal/repository"
	"livin-on-a-platter-api/internal/storage"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// YourHandler is the main handler for your route
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	repo := repository.NewImageUploadRepository()
	uuid := uuid.New()
	repo.Create(&img.ImageUpload{
		ID:        uuid,
		FileType:  "png",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	fmt.Println("message received")
}

func main() {
	firebase.GetDB().Connect()
	if err := storage.GetStorage().Connect(); err != nil {
		fmt.Printf("err during storage connection: %v\n", err)
	}
	http.HandleFunc("/huh", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hi there, I'm livin-on-a-platter-api!</h1>")
	})

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", api.Routes()); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}

}
