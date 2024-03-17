package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func Lightness() {
	file, err := os.Open("./Input/image.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err.Error())
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	Image := image.NewGray(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := (max(r, g, b) + min(r, g, b)) / 2
			//fmt.Println(r, " ", g, " ", b, " ", gray)
			Image.Set(x, y, color.Gray{Y: uint8(gray / 256)})
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

func Avarage() {
	file, err := os.Open("./Input/image.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err.Error())
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	Image := image.NewGray(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := (r + g + b) / 3
			//fmt.Println(r, " ", g, " ", b, " ", gray)
			Image.Set(x, y, color.Gray{Y: uint8(gray / 256)})
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

func Luminosity() {
	file, err := os.Open("./Input/image.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err.Error())
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	Image := image.NewGray(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
			//fmt.Println(r, " ", g, " ", b, " ", gray)
			Image.Set(x, y, color.Gray{Y: uint8(gray / 256)})
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
