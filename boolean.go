package main

import (
	"fmt"
	"image"
	"image/png"
	"os"
)

func invert() {
	file, err := os.Open("./Input/binerA.png")
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
	fmt.Println(image.Black)
	Image := image.NewRGBA64(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			fmt.Println(r, g, b)
			if r == 65535 && g == 65535 && b == 65535 {
				Image.Set(x, y, image.Black)
			} else {
				Image.Set(x, y, image.White)
			}

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

func AND() {
	file, err := os.Open("./Input/binerA.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	file2, err := os.Open("./Input/binerB.png")
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
			r, g, b, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 && r2 == 0 && g2 == 0 && b2 == 0 {
				Image.Set(x, y, image.Black)
			} else {
				Image.Set(x, y, image.White)
			}

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

func OR() {
	file, err := os.Open("./Input/binerA.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	file2, err := os.Open("./Input/binerB.png")
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
			r, g, b, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 || r2 == 0 && g2 == 0 && b2 == 0 {
				Image.Set(x, y, image.Black)
			} else {
				Image.Set(x, y, image.White)
			}

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

func XOR() {
	file, err := os.Open("./Input/binerA.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	file2, err := os.Open("./Input/binerB.png")
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
			r, g, b, _ := img.At(x, y).RGBA()
			r2, g2, b2, _ := img2.At(x, y).RGBA()
			if r == 0 && g == 0 && b == 0 && r2 == 65535 && g2 == 65535 && b2 == 65535 || r == 65535 && g == 65535 && b == 65535 && r2 == 0 && g2 == 0 && b2 == 0 {
				Image.Set(x, y, image.Black)
			} else {
				Image.Set(x, y, image.White)
			}

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
