package main

import (
	"fmt"
	_ "image/jpeg"
	"github.com/qeesung/image2ascii/convert"
)

func main(){
	convertOptions := convert.DefaultOptions
	convertOptions.FixedHeight = 50
	convertOptions.FixedWidth = 130

	converter := convert.NewImageConverter()
	fmt.Println(converter.ImageFile2ASCIIString("./download.jpeg", &convertOptions))
}