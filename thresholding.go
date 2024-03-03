package main

import (
	"image"
	"image/png"
	"os"
)

func thresholding() {
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

	Image := image.NewRGBA64(image.Rect(0, 0, width, height))
	min := 255
	max := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			sum := (r + g + b) / 3
			if max < int(sum) {
				max = int(sum)
			}
			if min > int(sum) {
				min = int(sum)
			}

		}
	}
	thr := (max + min) / 2
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			sum := (r + g + b) / 3
			if sum > uint32(thr) {
				Image.Set(x, y, image.White)
			} else {
				Image.Set(x, y, image.Black)

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
		return
	}
}
