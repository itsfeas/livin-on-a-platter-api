package middleware

import (
	"encoding/json"
	"net/http"
)

// SuccessResponseMiddleware is a middleware that sends a success JSON response for successful requests
func SuccessResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom ResponseWriter to intercept the status code
		customResponseWriter := &CustomResponseWriter{ResponseWriter: w}

		// Call the next handler in the chain
		next.ServeHTTP(customResponseWriter, r)

		// Check if the status code is within the success range (200-299)
		if customResponseWriter.statusCode >= http.StatusOK && customResponseWriter.statusCode < http.StatusMultipleChoices {
			// Create a generic success response
			successResponse := map[string]interface{}{
				"status": "ok",
			}

			// Convert the success response to JSON
			jsonResponse, err := json.Marshal(successResponse)
			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}

			// Set the content type to JSON and write the response
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
		}
	})
}

// CustomResponseWriter is a custom ResponseWriter that intercepts the status code
type CustomResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader intercepts the status code and stores it
func (w *CustomResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}
