package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func boxFilterSmoothing() {
	file, err := os.Open("./Input/Paimon.png")
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
	boxW := 1
	boxH := -1
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			var rTotal, gTotal, bTotal, count, Alpha uint32
			for dx := boxH; dx <= boxW; dx++ {
				for dy := boxH; dy <= boxW; dy++ {
					nx, ny := x+dx, y+dy
					if nx >= 0 && nx < width && ny >= 0 && ny < height {
						r, g, b, a := img.At(nx, ny).RGBA()
						rTotal += r
						gTotal += g
						bTotal += b
						Alpha = a
						count++
					}
				}
			}
			newImg.Set(x, y, color.RGBA{
				R: uint8((rTotal / count) / 256),
				G: uint8((gTotal / count) / 256),
				B: uint8((bTotal / count) / 256),
				A: uint8(Alpha),
			})
		}
	}

	outFile, err := os.Create("./Output/blur.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, newImg)
	if err != nil {
		panic(err)
	}
}
