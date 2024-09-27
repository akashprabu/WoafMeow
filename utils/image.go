package utils

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"

	"github.com/nfnt/resize"
)

// RotateImage rotates the given image and saves it to a new file
func RotateImage(inputPath, outputPath string) error {
	// Open the input image file
	imgFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Println("Error opening image file:", err)
		return err
	}
	defer imgFile.Close()

	// Decode the image
	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return err
	}

	// Rotate the image 90 degrees clockwise
	rotatedImg := rotate90Degrees(img)

	// Create the output file
	processedFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Println("Error creating processed image file:", err)
		return err
	}
	defer processedFile.Close()

	// Save the processed image as JPEG
	err = jpeg.Encode(processedFile, rotatedImg, nil)
	if err != nil {
		fmt.Println("Error encoding processed image:", err)
		return err
	}

	fmt.Println("Image successfully processed and saved as:", outputPath)
	return nil
}

// rotate90Degrees rotates the image 90 degrees clockwise
func rotate90Degrees(img image.Image) image.Image {
	bounds := img.Bounds()
	rotated := image.NewRGBA(image.Rect(0, 0, bounds.Dy(), bounds.Dx()))

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			rotated.Set(bounds.Dy()-y-1, x, img.At(x, y))
		}
	}

	return resize.Resize(500, 500, rotated, resize.Lanczos3) // Optionally resize
}
