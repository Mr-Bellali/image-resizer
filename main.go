package main

import (
    "errors"
    "flag"
    "fmt"
    "image"
    // "image/jpeg"
    "image/png"
    "os"
    "path/filepath"
    "strconv"
    "strings"

    "github.com/nfnt/resize"
)

// validateFile checks if a file exists and is of a valid image format (PNG or JPEG)
func validateFile(filePath string) error {
    file, err := os.Open(filePath)
    if err != nil {
        return errors.New("could not open file: please check the file path")
    }
    defer file.Close()

    // Check file format
    _, format, err := image.Decode(file)
    if err != nil {
        return errors.New("could not decode image: ensure the file is a valid image format (PNG or JPEG)")
    }
    if format != "png" && format != "jpeg" {
        return errors.New("unsupported file format: only PNG and JPEG are allowed")
    }
    return nil
}

// validateDimensions checks that the dimensions are valid positive integers
func validateDimensions(dimensions []string) ([]int, error) {
    var widths []int
    for _, dim := range dimensions {
        width, err := strconv.Atoi(dim)
        if err != nil || width <= 0 {
            return nil, fmt.Errorf("invalid dimension '%s': dimensions must be positive integers", dim)
        }
        widths = append(widths, width)
    }
    return widths, nil
}

func main() {
    // Define command-line flags
    fileFlag := flag.String("file", "", "Path to the original image file")
    sizesFlag := flag.String("sizes", "", "Space-separated list of sizes (e.g., '120 200 220')")
    outputDirFlag := flag.String("output", "images", "Directory to save resized images")

    // Parse flags
    flag.Parse()

    // Check if required flags are provided
    if *fileFlag == "" || *sizesFlag == "" {
        fmt.Println("Error: missing required flags. Use -file and -sizes.")
        fmt.Println("Usage: go run main.go -file <image_path> -sizes <size1 size2 ...> -output <output_dir>")
        return
    }

    // Validate input file
    if err := validateFile(*fileFlag); err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Validate dimensions
    dimensions := strings.Fields(*sizesFlag) // Split dimensions by spaces
    widths, err := validateDimensions(dimensions)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Create the output directory if it does not exist
    if _, err := os.Stat(*outputDirFlag); os.IsNotExist(err) {
        err := os.Mkdir(*outputDirFlag, 0755)
        if err != nil {
            fmt.Println("Error creating the output directory:", err)
            return
        }
    }

    // Open original image
    file, err := os.Open(*fileFlag)
    if err != nil {
        fmt.Println("Error opening the image:", err)
        return
    }
    defer file.Close()

    // Decode image
    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println("Error decoding the image:", err)
        return
    }

    // Loop through each width and resize the image
    for _, width := range widths {
        resizedImage := resize.Resize(uint(width), 0, img, resize.Lanczos3)

        // Save resized image
        outputFileName := filepath.Join(*outputDirFlag, fmt.Sprintf("%d.png", width))
        outputFile, err := os.Create(outputFileName)
        if err != nil {
            fmt.Println("Error creating the output file:", err)
            continue
        }

        err = png.Encode(outputFile, resizedImage)
        if err != nil {
            fmt.Println("Error encoding the resized image:", err)
            outputFile.Close()
            continue
        }

        outputFile.Close()
        fmt.Printf("Image saved as %s\n", outputFileName)
    }
}
