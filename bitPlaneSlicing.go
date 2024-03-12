package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
)

func bitPlane() {
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
	fmt.Sprintf("flag: %08b", flag)
}
