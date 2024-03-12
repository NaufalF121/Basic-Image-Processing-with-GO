package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func imageNegative() {

	file, err := os.Open("./Input/image2.png")
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
	var flag uint8 = 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.GrayModel.Convert(img.At(x, y))
			gray := c.(color.Gray)
			sum := gray.Y
			if sum > flag {
				flag = sum
			}
			Image.Set(x, y, gray)
		}
	}

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := Image.At(x, y)
			gray := c.(color.Gray)
			gray.Y = flag - gray.Y
			Image.Set(x, y, gray)
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
