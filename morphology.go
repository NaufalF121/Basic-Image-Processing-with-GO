package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func testImg() image.Image {
	width, height := 10, 10
	img := image.NewGray(image.Rect(0, 0, width, height))
	citra := [][]int{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if citra[x][y] == 1 {
				img.Set(x, y, color.Gray{Y: uint8(255)})
			} else {
				img.Set(x, y, color.Gray{Y: uint8(0)})
			}

		}
	}
	return img
}

func dilate(img image.Image) image.Image {

	radius := 1
	bounds := img.Bounds()
	newImg := image.NewGray(bounds)

	// Perform dilation
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			whiteFound := false

			// Check the surrounding pixels in a certain radius
			for i := -radius; i <= radius && !whiteFound; i++ {
				for j := -radius; j <= radius && !whiteFound; j++ {
					xn := x + i
					yn := y + j
					// If the pixel is out of bounds, skip it
					if xn < 0 || xn >= bounds.Dx() || yn < 0 || yn >= bounds.Dy() {
						continue
					}
					// If the pixel is white, set the current pixel to white
					gray := color.GrayModel.Convert(img.At(xn, yn)).(color.Gray).Y
					if gray == 255 {
						newImg.Set(x, y, color.Gray{Y: uint8(255)})
						whiteFound = true
					}
				}
			}
			// If none of the surrounding pixels is white, set the current pixel to black
			if !whiteFound {
				newImg.Set(x, y, color.Gray{Y: uint8(0)})
			}
		}
	}

	// Save the result
	outFile, err := os.Create("./Output/output.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, newImg)
	if err != nil {
		panic(err)
	}
	return newImg

}

func erode(img image.Image) image.Image {

	radius := 1

	bounds := img.Bounds()
	newImg := image.NewGray(bounds)

	// Perform erosion
	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			blackFound := false

			// Check the surrounding pixels in a certain radius
			for i := -radius; i <= radius && !blackFound; i++ {
				for j := -radius; j <= radius && !blackFound; j++ {
					xn := x + i
					yn := y + j
					// If the pixel is out of bounds, skip it
					if xn < 0 || xn >= bounds.Dx() || yn < 0 || yn >= bounds.Dy() {
						continue
					}
					// If the pixel is black, set the current pixel to black
					gray := color.GrayModel.Convert(img.At(xn, yn)).(color.Gray).Y
					if gray == 0 {
						newImg.Set(x, y, color.Gray{Y: uint8(0)})
						blackFound = true
					}
				}
			}
			// If none of the surrounding pixels is black, set the current pixel to white
			if !blackFound {
				newImg.Set(x, y, color.Gray{Y: uint8(255)})
			}
		}

	}
	file, err := os.Create("./Output/output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, newImg)
	if err != nil {
		panic(err)
	}
	return newImg
}

func opening() {
	img := testImg()
	out := dilate(erode(img))
	file, err := os.Create("./Output/output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, out)
	if err != nil {
		panic(err)
	}
}

func closing() {
	img := testImg()
	out := erode(dilate(img))
	file, err := os.Create("./Output/output.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, out)
	if err != nil {
		panic(err)
	}
}
