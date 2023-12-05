package http_util

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func WriteError(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("error Written: %w", error)

	message := map[string]string{
		"status": fmt.Sprint(code),
		"msg":    error,
		"time":   time.Now().String(),
	}
	jsonResponse, jsonErr := json.Marshal(message)
	if jsonErr != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		fmt.Println("Error encoding JSON in ErrorHandler")
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonResponse)
}
