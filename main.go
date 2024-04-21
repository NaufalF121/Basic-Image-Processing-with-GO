package main

const (
	input  = "./Input/image.png"
	output = "./Output/output"
	binerA = "./Input/binerA.png"
	binerB = "./Input/binerB.png"
)

type bitPlaneSlicing struct {
	input  string
	output string
	bitP   int
}

type boolean struct {
	input1 string
	input2 string
	output string
}

type Basic struct {
	input  string
	output string
}

type ColorModel struct {
	input    string
	channel1 bool // Red channel, Cyan Channel, Hue Channel, Y' Channel
	channel2 bool // Green channel, Magenta Channel, Saturation Channel, Cb Channel, U' Channel
	channel3 bool // Blue channel, Yellow Channel, Intensity Channel, Cr Channel, V' Channel
	output   string
}

func main() {
	// bitplane
	confBit := bitPlaneSlicing{
		input:  input,
		output: output + "bitplane.png",
		bitP:   5,
	}
	bitPlane(&confBit)

	// boolean operation
	confBool := boolean{
		input1: binerA,
		input2: binerB,
		output: output,
	}
	invert(&confBool) // in invert just take 1 input image so input1 is the one that used
	AND(&confBool)
	XOR(&confBool)
	OR(&confBool)

	// Color Model Operation
	confColor := ColorModel{
		input:    input,
		channel1: true,
		channel2: false,
		channel3: false,
		output:   output,
	}
	rgbToHsi(&confColor)
	rgbToYUV(&confColor)
	rgb2Cmyk(&confColor)
	rgbtoYCbCr(&confColor)

}
