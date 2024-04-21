package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func bitPlane(data *bitPlaneSlicing) {
	file, err := os.Open(data.input)
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
	// bitP := 1-9
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := Image.At(x, y)
			gray := c.(color.Gray)
			sum := fmt.Sprintf("%c", fmt.Sprintf("%08b", gray.Y)[int(8-data.bitP)])
			if sum == "1" {
				gray.Y = 255
			} else {
				gray.Y = 0
			}
			Image.Set(x, y, gray)
		}
	}
	outFile, err := os.Create(data.output)
	if err != nil {
		panic(err.Error())
	}
	defer outFile.Close()
	err = png.Encode(outFile, Image)
	if err != nil {
		panic(err.Error())
	}
}
