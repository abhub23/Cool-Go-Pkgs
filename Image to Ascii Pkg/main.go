package main

import (
	"fmt"
	_ "image/jpeg"
	"github.com/qeesung/image2ascii/convert"
)

func main(){
	convertOptions := convert.DefaultOptions
	convertOptions.FixedHeight = 50
	convertOptions.FixedWidth = 120

	converter := convert.NewImageConverter()
	fmt.Println(converter.ImageFile2ASCIIString("./download.jpeg", &convertOptions))
}

package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

// Function to decode JPEG from bytes
func decodeJPEG(imgData []byte) (image.Image, error) {
	reader := bytes.NewReader(imgData) // Convert []byte to io.Reader
	img, err := jpeg.Decode(reader)    // Decode JPEG
	if err != nil {
		return nil, err
	}
	return img, nil
}

func main() {
	// Step 1: Read the JPEG file into a byte slice
	imgBytes, err := os.ReadFile("download.jpeg") // Read entire file as []byte
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Step 2: Call decodeJPEG function with byte data
	img, err := decodeJPEG(imgBytes)
	if err != nil {
		fmt.Println("Error decoding JPEG:", err)
		return
	}

	// Step 3: Print image details (dimensions)
	fmt.Println("Decoded image bounds:", img.Bounds()) // Example output: (0,0)-(1920,1080)
	imgSize := float64 (len(imgBytes)) / 1000
	fmt.Printf("The image is %f KB",imgSize)
}
