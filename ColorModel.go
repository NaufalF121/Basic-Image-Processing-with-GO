package main

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"

	"Basic__Image_Processing/Colorutil"
	_ "Basic__Image_Processing/Colorutil"
)

// HSI is a color model that is used to represent the color of an image. HSI stands for Hue, Saturation, and Intensity.
func rgbToHsi(data *ColorModel) {
	file, err := os.Open(data.input)
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
	hsiImg := Colorutil.NewNHSVA(bounds)

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
			if !data.channel1 {
				h = 0
			}
			if !data.channel2 {
				s = 0
			}
			if !data.channel3 {
				i = 0
			}
			hsiImg.SetNHSVA(x, y, Colorutil.NHSIA{
				H: uint8(h * 255.0 / (2 * math.Pi)),
				S: uint8(s * 255.0),
				I: uint8(i * 255.0),
				A: 255,
			})
		}
	}

	outFile, err := os.Create(data.output + "hsi_image.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, hsiImg)
	if err != nil {
		panic(err)
	}
}

// YUV is a color model that is used to represent the color of an image. YUV stands for Luminance (Y), Chrominance Blue (U), and Chrominance Red (V).
func rgbToYUV(data *ColorModel) {
	file, err := os.Open(data.input)
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
	hsiImg := Colorutil.NewYUVA(bounds)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			rFloat := float64(r) / 65535.0
			gFloat := float64(g) / 65535.0
			bFloat := float64(b) / 65535.0

			Y := 0.257*rFloat + 0.504*gFloat + 0.098*bFloat
			v := (0.439*rFloat - 0.368*gFloat + 0.071*bFloat) + 128
			u := (0.148*rFloat - 0.291*gFloat + 0.439*bFloat) + 128
			if !data.channel1 {
				Y = 0
			}
			if !data.channel2 {
				u = 0
			}
			if !data.channel3 {
				v = 0
			}
			hsiImg.SetYUVA(x, y, Colorutil.YUV{
				Y: uint8(Y * 255.0),
				U: uint8(u * 255.0),
				V: uint8(v * 255.0),
				A: 255,
			})
		}
	}

	outFile, err := os.Create(data.output + " yuv_image.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, hsiImg)
	if err != nil {
		panic(err)
	}
}

// CMYK is a color model that is used to represent the color of an image. CMYK stands for Cyan, Magenta, Yellow, and Key (Black).
func rgb2Cmyk(data *ColorModel) {
	file, err := os.Open(data.input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	cmykImg := image.NewCMYK(bounds)

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()

			rFloat := float64(r) / 65535.0
			gFloat := float64(g) / 65535.0
			bFloat := float64(b) / 65535.0

			k := 1 - max(max(rFloat, gFloat), bFloat)
			c := (1 - rFloat - k) / (1 - k)
			m := (1 - gFloat - k) / (1 - k)
			Y := (1 - bFloat - k) / (1 - k)

			if !data.channel1 {
				c = 0
			}
			if !data.channel2 {
				m = 0
			}
			if !data.channel3 {
				Y = 0
			}

			cmykImg.Set(x, y, color.CMYK{
				C: uint8(c * 255.0),
				M: uint8(m * 255.0),
				Y: uint8(Y * 255.0),
				K: uint8(k * 255.0),
			})
		}
	}

	outFile, err := os.Create(data.output + " cmyk_image.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	err = png.Encode(outFile, cmykImg)
	if err != nil {
		panic(err)
	}

}

// YCbCr is a color model that is used to represent the color of an image. YCbCr stands for Luminance (Y), Chrominance Blue (Cb), and Chrominance Red (Cr).
func rgbtoYCbCr(data *ColorModel) {
	file, err := os.Open(data.input)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, err := png.Decode(file)

	if err != nil {
		panic(err)
	}
	bounds := img.Bounds()
	ycbcrImg := image.NewYCbCr(bounds, image.YCbCrSubsampleRatio444)

	for x := 0; x < bounds.Dx(); x++ {
		for y := 0; y < bounds.Dy(); y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			Y, Cb, Cr := color.RGBToYCbCr(uint8(r>>8), uint8(g>>8), uint8(b>>8))
			i := ycbcrImg.YOffset(x, y)
			if !data.channel1 {
				Y = 0
			}
			if !data.channel2 {
				Cb = 0
			}
			if !data.channel3 {
				Cr = 0
			}
			ycbcrImg.Y[i] = Y
			ycbcrImg.Cb[i] = Cb
			ycbcrImg.Cr[i] = Cr
		}
	}
	file, err = os.Create(data.output + " ycbcr_image.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	err = png.Encode(file, ycbcrImg)
	if err != nil {
		panic(err)
	}
}
