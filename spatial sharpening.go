package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

var laplaceFilter = [3][3]int{
	{-1, -1, -1},
	{-1, 8, -1},
	{-1, -1, -1},
}

func applyLaplaceFilter() {
	file, err := os.Open("./Output/blur.png")
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
			var rTotal, gTotal, bTotal int32
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx, ny := x+dx, y+dy
					if nx >= 0 && nx < width && ny >= 0 && ny < height {
						r, g, b, _ := img.At(nx, ny).RGBA()
						weight := laplaceFilter[dx+1][dy+1]
						rTotal += int32(r) * int32(weight)
						gTotal += int32(g) * int32(weight)
						bTotal += int32(b) * int32(weight)
					}
				}
			}
			rTotal = max(0, min(65535, rTotal))
			gTotal = max(0, min(65535, gTotal))
			bTotal = max(0, min(65535, bTotal))
			newImg.Set(x, y, color.RGBA64{
				R: uint16(rTotal),
				G: uint16(gTotal),
				B: uint16(bTotal),
				A: 65535,
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

func addition() {
	file, err := os.Open("./Output/blur.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	lap, err := os.Open("./Output/output.png")
	if err != nil {
		panic(err)
	}
	defer lap.Close()
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	img2, err := png.Decode(lap)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	newImg := image.NewRGBA(bounds)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			r2, g2, b2, a2 := img2.At(x, y).RGBA()
			newImg.Set(x, y, color.RGBA{R: uint8(min(255, r>>8+r2>>8)), G: uint8(min(255, g>>8+g2>>8)), B: uint8(min(255, b>>8+b2>>8)), A: uint8(min(255, a>>8+a2>>8))})
		}
	}
	outFile, err := os.Create("./Output/output1.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	err = png.Encode(outFile, newImg)
	if err != nil {
		panic(err)
	}
}
