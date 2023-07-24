package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

func convertToWebP(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isImageFile(path) {
			fmt.Printf("Converting %s to WebP...\n", path)

			// Open the image file
			f, err := os.Open(path)
			if err != nil {
				return err
			}
			defer f.Close()

			// Decode the image
			img, err := imaging.Decode(f)
			if err != nil {
				return err
			}

			// Create a new file with .webp extension
			newFilename := path[:len(path)-len(filepath.Ext(path))] + ".webp"
			newFile, err := os.Create(newFilename)
			if err != nil {
				return err
			}
			defer newFile.Close()

			// Encode the image as WebP and write it to the new file
			err = webp.Encode(newFile, img, nil)
			if err != nil {
				return err
			}

			// Delete the original image file
			err = os.Remove(path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func isImageFile(filename string) bool {
	ext := filepath.Ext(filename)
	switch ext {
	case ".jpg", ".jpeg", ".png", ".gif":
		return true
	}
	return false
}

func main() {
	absPath, err := filepath.Abs("/home/hungptran/hungon.space/resources/_gen/images")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = convertToWebP(absPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
