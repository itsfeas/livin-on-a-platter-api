package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error (in a real-world application, you might want to log to a file or a service)
				fmt.Printf("Internal Error: %v\n", err)

				// Respond with an Internal Server Error
				w.Header().Set("Content-Type", "application/json")
				http.Error(w, "Error on server", http.StatusInternalServerError)

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
		}()

		// Let the next handler in the chain handle the request
		next.ServeHTTP(w, r)
	})
}
