package middleware

import (
	"fmt"
	response "livin-on-a-platter-api/internal/responses/types"
	"net/http"
)

// ErrorHandlerMiddleware is a middleware that handles internal server errors
func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Log the error (in a real-world application, you might want to log to a file or a service)
				fmt.Printf("Panic: %v\n", err)

				// Respond with an Internal Server Error
				w.Header().Set("Content-Type", "application/json")

				resp := &response.BaseMsg{
					Status: http.StatusInternalServerError,
					Msg:    "Error on server",
				}

				jsonResponse, jsonErr := resp.ToJson()
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
