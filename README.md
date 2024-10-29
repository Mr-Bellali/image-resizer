
# Image Resizer

This is a simple command-line tool to resize an image to multiple specified dimensions.

## Features

- Resizes an image to multiple specified widths
- Saves the resized images in a specified output directory
- Supports PNG and JPEG formats

## Installation

1. Make sure you have Go installed. If not, download it from [Go's official website](https://golang.org/).
2. Clone the repository:
   ```bash
   git clone https://github.com/Mr-Bellali/image-resizer
   ```
3. Build the executable:
   ```bash
   go build -o resizer main.go
   ```

## Usage

```bash
go run main.go -file <path_to_image> -sizes "<size1 size2 size3 ...>" -output <output_directory>
```

### Example

Resize `logo.png` to widths of 120, 200, and 300 pixels, and save the output in the `images` directory:

```bash
go run main.go -file logo.png -sizes "120 200 300" -output images
```

### Flags

- `-file`: Path to the original image file (required)
- `-sizes`: Space-separated list of widths to resize the image (required)
- `-output`: Directory to save resized images (default: `images`)

## Error Handling

- If an image file does not exist or is an unsupported format, the program will print a helpful message.
- Invalid dimensions will be skipped with an error message.


