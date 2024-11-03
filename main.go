package main

import (
	"bytes"
	"fmt"
	"image"
	"image/png" // Import for PNG format support
	"log"
	"os"

	// Required for embedding
	_ "embed"

	"github.com/giant-stone/goimg"
)

//go:embed assets/ad.png
var exampleImage []byte

// var characterSet = map[uint8]rune{
// 	0: ' ',
// 	1: '`',
// 	2: '.',
// 	3: '~',
// 	4: '^',
// 	5: '=',
// 	6: '0',
// 	7: '@',
// }

var characterSet = map[uint8]rune{
	0: ' ',
	1: '⠠',
	2: '⠢',
	3: '⠣',
	4: '⠵',
	5: '⠽',
	6: '⠿',
	7: '⠿',
}

func main() {
	// Load the image from embedded bytes
	img, _, err := image.Decode(bytes.NewReader(exampleImage))
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	aspectRatio := float64(img.Bounds().Max.X) / float64(img.Bounds().Max.Y)
	resizedY := 376 / aspectRatio
	fmt.Println("y", resizedY)

	resizedImg := goimg.Resize(376, uint(resizedY), img, goimg.Lanczos3)

	output, err := os.Create("smol.png")
	if err != nil {
		log.Fatal(err)
	}

	defer output.Close()

	err = png.Encode(output, resizedImg)
	if err != nil {
		log.Fatal(err)
	}

	grayImg := imgToGray(resizedImg)

	for i := range grayImg {
		for j := range grayImg[i] {
			fmt.Printf("%c", minGrayToAscii(grayToMinGray(grayImg[i][j])))
		}
		fmt.Println()
	}

}

func grayToMinGray(x uint8) uint8 {
	return x / 32 // 8 characterSet
}

func imgToGray(img image.Image) [][]uint8 {
	newImage := make([][]uint8, img.Bounds().Max.Y)
	for y := 0; y < img.Bounds().Max.Y; y++ {
		newImage[y] = make([]uint8, img.Bounds().Max.X)
		for x := 0; x < img.Bounds().Max.X; x++ {
			var rgba = img.At(x, y)
			r, g, b, _ := rgba.RGBA()
			//gray := uint8(0.33*float64(r) + 0.33*float64(g) + 0.33*float64(b))
			gray := uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
			newImage[y][x] = gray
		}
	}
	return newImage
}

func minGrayToAscii(pixel uint8) rune {
	return characterSet[pixel]
}
