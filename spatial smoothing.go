package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func boxFilterSmoothing() {
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
	newImg := image.NewRGBA(bounds)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var rTotal, gTotal, bTotal, count uint32
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx, ny := x+dx, y+dy
					if nx >= 0 && nx < width && ny >= 0 && ny < height {
						r, g, b, _ := img.At(nx, ny).RGBA()
						rTotal += r
						gTotal += g
						bTotal += b
						count++
					}
				}
			}
			newImg.Set(x, y, color.RGBA{
				R: uint8(rTotal / count >> 8),
				G: uint8(gTotal / count >> 8),
				B: uint8(bTotal / count >> 8),
				A: 255,
			})
		}
	}

	outFile, err := os.Create("./Output/output.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, newImg)
	if err != nil {
		panic(err)
	}
}
