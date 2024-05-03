package main

import (
	"fmt"
	"gif_test_mods/folderopener"
	"image/gif"
	"os"
	"path/filepath"
)

func main() {
	folderPath, value1, value2, value3 := folderopener.OpenFolder()
	fmt.Println("Selected folder:", folderPath)
	outFolderPath := "output_directory"

	if _, err := os.Stat(outFolderPath); err == nil {
		if err := os.RemoveAll(outFolderPath); err != nil {
			panic(err) // Handle the error if removal fails
		}
	}

	// Create the directory with 0755 permissions
	if err := os.Mkdir(outFolderPath, 0755); err != nil {
		panic(err) // Handle the error if creation fails
	}

	inputDir := folderPath          // Change this to your input directory path
	outputDir := "output_directory" // Change this to your output directory path
	fmt.Printf("value %d", value1)
	fmt.Printf("value 2 %d", value2)
	fmt.Printf("value 3 %d", value3)
	// Open the input directory
	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("Error accessing path %s: %v\n", path, err)
			return err
		}

		// Check if the file is a regular file and has a .gif extension
		if !info.IsDir() && filepath.Ext(path) == ".gif" {
			// Open the GIF file
			file, err := os.Open(path)
			if err != nil {
				fmt.Printf("Error opening file %s: %v\n", path, err)
				return nil
			}
			defer file.Close()

			// Decode the GIF file
			gifImg, err := gif.DecodeAll(file)
			if err != nil {
				fmt.Printf("Error decoding GIF %s: %v\n", path, err)
				return nil
			}

			// Optimize the GIF
			optimizedGIF := optimizeGIF(gifImg)

			// Create the output directory if it doesn't exist
			err = os.MkdirAll(outputDir, os.ModePerm)
			if err != nil {
				fmt.Printf("Error creating output directory %s: %v\n", outputDir, err)
				return nil
			}

			// Write the optimized GIF to a new file in the output directory
			outputFile, err := os.Create(filepath.Join(outputDir, info.Name()))
			if err != nil {
				fmt.Printf("Error creating output file for %s: %v\n", path, err)
				return nil
			}
			defer outputFile.Close()

			// Encode the optimized GIF and write it to the output file
			err = gif.EncodeAll(outputFile, optimizedGIF)
			if err != nil {
				fmt.Printf("Error encoding GIF %s: %v\n", path, err)
				return nil
			}

			fmt.Printf("GIF optimization completed for %s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error accessing directory: %v\n", err)
	}
}

func optimizeGIF(g *gif.GIF) *gif.GIF {
	// Check if the number of frames exceeds the limit (999)
	if len(g.Image) > 999 {
		// Trim frames to the first 999 frames
		g.Image = g.Image[:999]
		g.Delay = g.Delay[:999]
		g.Disposal = g.Disposal[:999]
	}

	return g
}
