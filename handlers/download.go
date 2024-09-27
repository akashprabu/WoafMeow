package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

// DownloadImageHandler serves the processed image for download
func DownloadImageHandler(w http.ResponseWriter, r *http.Request) {
	fileName := filepath.Base(r.URL.Path[len("/download/"):])
	processedPath := filepath.Join("./processed", fileName)

	if _, err := os.Stat(processedPath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "image/jpeg")

	http.ServeFile(w, r, processedPath)
}
