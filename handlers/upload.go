package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"WoafMeow/utils"
)

// UploadImageHandler handles image uploads
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // 10MB limit
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Failed to retrieve image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a unique file name for the uploaded image
	fileName := fmt.Sprintf("uploaded_image_%d.jpg", time.Now().UnixNano())
	filePath := filepath.Join("./uploads", fileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to save image", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Failed to write image to file", http.StatusInternalServerError)
		return
	}

	processedPath := filepath.Join("./processed", fileName)
	err = utils.RotateImage(filePath, processedPath)
	if err != nil {
		http.Error(w, "Failed to process image", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(fmt.Sprintf("Image uploaded and processed successfully as: %s", fileName)))
}
