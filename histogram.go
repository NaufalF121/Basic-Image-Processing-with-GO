package main

import (
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

// generate random data for bar chart
func generateBarItems2() []opts.BarData {
	items := make([]opts.BarData, 256)
	for i := 0; i < 256; i++ {
		items[i] = opts.BarData{Value: 0}
	}
	file, err := os.Open("./Output/output.png")
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
	//Image := image.NewGray(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
			colour := color.Gray{Y: uint8(gray / 256)}

			items[(colour.Y)].Value = items[colour.Y].Value.(int) + 1
			//Image.Set(x, y, gray)
		}
	}
	return items
}

// generate random data for bar chart
func generateBarItems() []opts.BarData {
	items := make([]opts.BarData, 256)
	for i := 0; i < 256; i++ {
		items[i] = opts.BarData{Value: 0}
	}
	file, err := os.Open("./Input/img.png")
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
	//Image := image.NewGray(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
			colour := color.Gray{Y: uint8(gray / 256)}

			items[(colour.Y)].Value = items[colour.Y].Value.(int) + 1
			//Image.Set(x, y, gray)
		}
	}
	return items
}
func equalize() []opts.BarData {
	file, err := os.Open("./Input/img.png")
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
	items := generateBarItems()
	for i := 1; i < 256; i++ {
		items[i] = opts.BarData{Value: items[i].Value.(int) + items[i-1].Value.(int)}
	}
	for i := 0; i < 256; i++ {
		items[i] = opts.BarData{Value: uint8(255 * (float64(items[i].Value.(int)) / float64(width*height)))}
		fmt.Println(items[i].Value)
	}
	Image := image.NewGray(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			gray := 0.21*float64(r) + 0.72*float64(g) + 0.07*float64(b)
			colour := color.Gray{Y: uint8(gray / 256)}
			colour.Y = items[colour.Y].Value.(uint8)
			Image.Set(x, y, colour)
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
	return items
}
func createBarChart() {
	// create a new bar instance
	bar := charts.NewBar()

	// Set global options
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Histogram",
		Subtitle: "Original Image VS Equalized Image",
	}))
	label := []string{}
	for i := 0; i < 256; i++ {
		label = append(label, strconv.Itoa(i))
	}
	equalize()
	//fmt.Println(generateBarItems())
	// Put data into instance
	bar.SetXAxis(label).
		AddSeries("Original", generateBarItems()).AddSeries("Equalized", generateBarItems2())
	f, _ := os.Create("bar.html")
	_ = bar.Render(f)
}
