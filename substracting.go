package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func substraction() {
	file, err := os.Open("./Output/blur.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	file2, err := os.Open("./Output/output.png")
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
	Image := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			Image.Set(x, y, color.RGBA{R: uint8(max(0, r>>8-r2>>8)), G: uint8(max(0, g>>8-g2>>8)), B: uint8(max(0, b>>8-b2>>8)), A: uint8(a >> 8)})
		}
	}
	outFile, err := os.Create("./Output/output1.png")
	if err != nil {
		panic(err.Error())
	}
	defer outFile.Close()
	err = png.Encode(outFile, Image)
	if err != nil {
		panic(err.Error())
	}
}
