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

    // Check if the file exists
    if _, err := os.Stat(processedPath); os.IsNotExist(err) {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }

    // Serve the file
    http.ServeFile(w, r, processedPath)
}
