package middleware

import (
	"encoding/json"
	"net/http"
)

func SuccessResponse(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		message := map[string]string{"status": "ok"}
		jsonResponse, err := json.Marshal(message)

		if err != nil {
			return err
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
		return nil
	})
}
