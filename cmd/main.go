package main

import (
	"fmt"
	"livin-on-a-platter-api/internal/db/firebase"
	"livin-on-a-platter-api/internal/middleware"
	"livin-on-a-platter-api/internal/model"
	"livin-on-a-platter-api/internal/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// YourHandler is the main handler for your route
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	repo := repository.NewImageRepository()
	repo.Create(&model.ImageUpload{
		ID:        uuid.New(),
		FileName:  "file",
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	fmt.Println("message received")
}

func main() {
	firebase.GetDB().Connect()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hi there, I'm livin-on-a-platter-api!</h1>")
	})

	mainHandler := http.HandlerFunc(YourHandler)
	manager := middleware.MiddlewareManager(mainHandler, middleware.SuccessResponseMiddleware, middleware.ErrorHandler)

	fmt.Println("Server is listening on port 8080...")
	if err := http.ListenAndServe(":8080", manager); err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}

}
