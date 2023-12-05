package middleware

import (
	"encoding/json"
	"fmt"
	"livin-on-a-platter-api/internal/middleware/writer"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a custom ResponseWriter to intercept the status code
		writer := &writer.CustomResponseWriter{ResponseWriter: w}

		// Let the next handler in the chain handle the request
		next.ServeHTTP(w, r)

		// Check if the status code is within the success range (200-299)
		if writer.StatusCode < http.StatusOK || writer.StatusCode >= http.StatusMultipleChoices {
			// Log the error (in a real-world application, you might want to log to a file or a service)
			fmt.Printf("Internal Error: %v\n", writer.Msg)
			w.Write([]byte{})

			// Respond with an Internal Server Error
			w.Header().Set("Content-Type", "application/json")

			message := map[string]string{"status": "error"}
			jsonResponse, jsonErr := json.Marshal(message)
			if jsonErr != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				fmt.Println("Error encoding JSON in ErrorHandler")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			w.Write(jsonResponse)
		}
	})
}
