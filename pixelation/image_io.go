package pixelation

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"path/filepath"
)

func OpenImage(imagePath string) image.Image {
	file, _ := os.Open(imagePath)
	defer file.Close()
	img, _, _ := image.Decode(file)
	return img
}

func SaveImage(imagePath string, img image.Image) error {
	// return extension of this file name
	ext := filepath.Ext(imagePath)
	dir := filepath.Dir(imagePath)
	newImagePath := fmt.Sprintf("%s/result%s", dir, ext)
	file, _ := os.Create(newImagePath)
	defer file.Close()
	err := jpeg.Encode(file, img, nil)
	if err != nil {
		return err
	}
	return nil
}
