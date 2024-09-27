package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"WoafMeow/handlers"
)

const (
	uploadDir    = "./uploads"
	processedDir = "./processed"
)

func main() {
	// Create directories if they don't exist
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(processedDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Howdy!, Welcome to WoafMeow Image Processor!")
	http.HandleFunc("/upload", handlers.UploadImageHandler)
	http.HandleFunc("/download/", handlers.DownloadImageHandler)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
