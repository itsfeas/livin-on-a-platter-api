package api

import (
	upload_routes "livin-on-a-platter-api/api/v1/routes/uploads"
	view_routes "livin-on-a-platter-api/api/v1/routes/view"
	"livin-on-a-platter-api/internal/middleware"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", upload_routes.UploadRoutes()))
	mux.Handle("/api/v1/view/", http.StripPrefix("/api/v1/view", view_routes.ViewRoutes()))
	wrappedRoutes := middleware.MiddlewareManager(mux, middleware.CorsManagerMiddleware, middleware.ErrorHandlerMiddleware)
	return wrappedRoutes
}
