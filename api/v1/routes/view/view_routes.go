package view_routes

import (
	handlers "livin-on-a-platter-api/api/v1/handlers/view"
	"net/http"
)

func ViewRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.ViewHandler)
	return mux
}
