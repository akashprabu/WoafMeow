# WoafMeow Image Processor API

## Overview
WoafMeow Image Processor is a simple API service that allows users to upload an image, processes it by rotating 90 degree, and then allows users to download the processed image. The API is built using Go and is structured in a modular way for better organization and maintainability.

## Folder Structure
```
WoafMeow/
├── main.go                # Entry point of the application
├── handlers/              # Contains HTTP handlers for API routes
│   ├── upload.go          # Handles image upload functionality
│   └── download.go        # Handles image download functionality
├── utils/                 # Utility functions for image processing
│   └── image.go           # Contains functions for image manipulation
├── uploads/               # Directory for storing uploaded images
├── processed/             # Directory for storing processed images
└── README.md              # Project documentation
```

## Requirements
- Go (version 1.18 or higher)
- Git (for version control)

## Installation

1. **Clone the Repository**
   ```bash
   git clone https://github.com/akashprabu/WoafMeow.git
   cd WoafMeow
   ```

2. **Install Go Modules**
   Make sure you are in the root directory of the project and run:
   ```bash
   go get -u github.com/nfnt/resize
   go mod tidy
   ```

3. **Run the Application**
   Start the server by executing:
   ```bash
   go run main.go
   ```

   The server will start on `http://localhost:8080`.

## API Endpoints

### 1. Upload an Image
- **Endpoint**: `/upload`
- **Method**: `POST`
- **Request**:
  ```bash
  curl -X POST -F "image=@/path/to/your/image.jpg" http://localhost:8080/upload
  ```
- **Response**: A message indicating whether the upload was successful.

### 2. Download a Processed Image
- **Endpoint**: `/download/{filename}`
- **Method**: `GET`
- **Request**:
  ```bash
  curl -O http://localhost:8080/download/processed_image_name.jpg
  ```
- **Response**: The processed image file will be downloaded.

## Notes
- Uploaded images are stored in the `uploads` directory.
- Processed images (rotated images) are stored in the `processed` directory.
- Make sure to check file types; supported formats include `.jpg`, `.jpeg`, and `.png`.

## Troubleshooting
If you encounter issues, consider the following:
- Ensure the Go environment is properly set up.
- Check for required permissions for reading/writing files in the `uploads` and `processed` directories.
- Verify the URL and the image file path when using `curl`.