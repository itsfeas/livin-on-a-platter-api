package handlers

import (
	"errors"
	"fmt"
	"livin-on-a-platter-api/internal/db/storage"
	msg "livin-on-a-platter-api/internal/model/msg/types"
	http_util "livin-on-a-platter-api/internal/util/error"
	"net/http"
	"os"
	"regexp"
)

var uuidRegex = regexp.MustCompile(`/([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})`)

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	bucket := os.Getenv("IMG_STORAGE_BUCKET")
	fmt.Println("bucket: ", bucket)

	fb := storage.GetStorage()

	imgId, err := getImgIdFromUrl(r.URL.Path)
	if err != nil {
		http_util.WriteError(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	imgUrl, err := fb.ExposeImage(imgId, bucket)
	if err != nil {
		http_util.WriteError(w, "Could not expose image URL", http.StatusInternalServerError)
		return
	}

	respond(w, imgUrl)
}

func getImgIdFromUrl(url string) (string, error) {
	match := uuidRegex.FindStringSubmatch(url)
	if len(match) < 2 {
		return "", errors.New("invalid URL")
	}
	return match[1], nil
}

func respond(w http.ResponseWriter, imgUrl string) {
	// Create a generic success response
	resp := msg.DefaultDataMsg()
	resp.Data = map[string]interface{}{
		"imgUrl": imgUrl,
	}

	// Convert the success response to JSON
	jsonResp, err := resp.ToJson()
	if err != nil {
		http_util.WriteError(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(jsonResp); err != nil {
		http_util.WriteError(w, "Error writing JSON", http.StatusInternalServerError)
		return
	}
}
