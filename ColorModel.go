package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

func rgbToHsi() {
	file, err := os.Open("./Input/image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	hsiImg := image.NewRGBA(bounds)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			rFloat := float64(r) / 65535.0
			gFloat := float64(g) / 65535.0
			bFloat := float64(b) / 65535.0

			i := (rFloat + gFloat + bFloat) / 3.0

			minVal := math.Min(math.Min(rFloat, gFloat), bFloat)
			s := 1 - 3.0*(minVal/(rFloat+gFloat+bFloat))

			var h float64
			if s == 0 {
				h = 0
			} else {
				numerator := ((rFloat - gFloat) + (rFloat - bFloat)) / 2
				denominator := math.Sqrt((rFloat-gFloat)*(rFloat-gFloat) + (rFloat-bFloat)*(gFloat-bFloat))
				theta := math.Acos(numerator / denominator)

				if bFloat <= gFloat {
					h = theta
				} else {
					h = 2*math.Pi - theta
				}
			}

			hsiImg.Set(x, y, color.RGBA{
				R: uint8(h * 255.0 / (2 * math.Pi)),
				G: uint8(s * 255.0),
				B: uint8(i * 255.0),
				A: 255,
			})
		}
	}

	outFile, err := os.Create("./Output/hsi_image.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, hsiImg)
	if err != nil {
		panic(err)
	}
}
