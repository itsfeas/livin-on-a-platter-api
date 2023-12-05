package routes

import (
	handlers "livin-on-a-platter-api/api/v1/handlers/uploads"
	"net/http"
)

func UploadRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/upload", handlers.UploadHandler)
	return mux
}
