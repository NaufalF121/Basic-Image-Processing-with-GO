package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func substraction() {
	file, err := os.Open("./Input/image.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	file2, err := os.Open("./Input/image2.png")
	if err != nil {
		panic(err.Error())
	}
	defer file2.Close()
	img, err := png.Decode(file)
	if err != nil {
		panic(err.Error())
	}
	img2, err := png.Decode(file2)
	if err != nil {
		panic(err.Error())
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	Image := image.NewRGBA64(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			Image.Set(x, y, color.RGBA64{R: uint16(r - r2), G: uint16(g - g2), B: uint16(b - b2), A: uint16(a)})
		}
	}
	outFile, err := os.Create("./Output/output.png")
	if err != nil {
		panic(err.Error())
	}
	defer outFile.Close()
	err = png.Encode(outFile, Image)
	if err != nil {
		panic(err.Error())
	}
}
