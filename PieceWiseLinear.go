package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func ContrastStretching() {
	file, err := os.Open("./Input/Low.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		panic(err.Error())
	}

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	Image := image.NewGray(image.Rect(0, 0, width, height))
	var flagHigh uint8 = 0
	var flagLow uint8 = 255
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := color.GrayModel.Convert(img.At(x, y))
			gray := c.(color.Gray)
			sum := gray.Y
			if sum > flagHigh {
				flagHigh = sum
			}
			if sum < flagLow {
				flagLow = sum
			}
			Image.Set(x, y, gray)
		}
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			c := Image.At(x, y)
			gray := c.(color.Gray)
			//Io = (Ii-Mini)*(((Maxo-Mino)/(Maxi-Mini))+Mino)
			//Io                                - Output pixel value
			//Ii                                 - Input pixel value
			//Mini                         - Minimum pixel value in the input image
			//Maxi                        - Maximum pixel value in the input image
			//Mino                        - Minimum pixel value in the output image (0)
			//Maxo                       - Maximum pixel value in the output image (255)
			gray.Y = uint8(gray.Y-flagLow) * (255 / (flagHigh - flagLow))
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
