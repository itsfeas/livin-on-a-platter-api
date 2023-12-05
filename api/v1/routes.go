package api

import (
	routes "livin-on-a-platter-api/api/v1/routes/uploads"
	"livin-on-a-platter-api/internal/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", routes.UploadRoutes()))
	wrappedRoutes := middleware.MiddlewareManager(mux, middleware.CorsManagerMiddleware, middleware.ErrorHandlerMiddleware)
	return wrappedRoutes
}
