package main

import (
	"fmt"
	"github.com/spakin/hsvimage"
	"github.com/spakin/hsvimage/hsvcolor"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func multiplication() {
	file, err := os.Open("./Input/image.png")
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err.Error())
	}
	skalar := 2

	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	Image := hsvimage.NewNHSVAF64(image.Rect(0, 0, width, height))
	draw.Draw(Image, Image.Bounds(), img, bounds.Min, draw.Src)
	fmt.Println(img)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {

			h, g, b, a := Image.NHSVAF64At(x, y).H, Image.NHSVAF64At(x, y).S, Image.NHSVAF64At(x, y).V, Image.NHSVAF64At(x, y).A

			Image.SetNHSVAF64(x, y, hsvcolor.NHSVAF64{H: h, S: g, V: b * float64(skalar), A: a})
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
