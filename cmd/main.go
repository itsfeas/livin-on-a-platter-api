package cmd

import (
	"fmt"
	"net/http"
	"livin-on-a-platter-api/internal/middleware"
)

// YourHandler is the main handler for your route
func YourHandler(w http.ResponseWriter, r *http.Request) {
	// Your logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, Middleware!"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hi there, I'm livin-on-a-platter-api!</h1>")
	})

	handler := http.HandleFunc("/api/", YourHandler)
	http.HandleFunc(middleware.MiddlewareManager(
		handler, middleware.SuccessResponseMiddleware, middleware.ErrorHandler)
	)

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)

}
