package Colorutil

import (
	"image"
	"image/color"
)

type NHSIAImage struct {
	// Pix holds the image's pixels, in H, S, V, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *NHSIAImage) ColorModel() color.Model { return NHSIAModel }

// Bounds returns the image's bounding rectangle.
func (p *NHSIAImage) Bounds() image.Rectangle { return p.Rect }

// At returns the color at the given image coordinates.
func (p *NHSIAImage) At(x, y int) color.Color {
	return p.NHSIAAt(x, y)
}

func (p *NHSIAImage) NHSIAAt(x, y int) NHSIA {
	if !(image.Point{x, y}.In(p.Rect)) {
		return NHSIA{}
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	return NHSIA{H: s[0], S: s[1], I: s[2], A: s[3]}
}
func (p *NHSIAImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
}

func (p *NHSIAImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := NHSIAModel.Convert(c).(NHSIA)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = c1.H
	s[1] = c1.S
	s[2] = c1.I
	s[3] = c1.A
}

func (p *NHSIAImage) SetNHSVA(x, y int, c NHSIA) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = c.H
	s[1] = c.S
	s[2] = c.I
	s[3] = c.A
}

func (p *NHSIAImage) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to
	// be inside either r1 or r2 if the intersection is empty. Without
	// explicitly checking for this, the Pix[i:] expression below can
	// panic.
	if r.Empty() {
		return &NHSIAImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &NHSIAImage{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

func NewNHSVA(r image.Rectangle) *NHSIAImage {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 4*w*h)
	return &NHSIAImage{pix, 4 * w, r}
}

type YUVAImage struct {
	// Pix holds the image's pixels, in H, S, V, A order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*4].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func (p *YUVAImage) ColorModel() color.Model { return YUVModel }

// Bounds returns the image's bounding rectangle.
func (p *YUVAImage) Bounds() image.Rectangle { return p.Rect }

// At returns the color at the given image coordinates.
func (p *YUVAImage) At(x, y int) color.Color {
	return p.YUVAAt(x, y)
}

func (p *YUVAImage) YUVAAt(x, y int) YUV {
	if !(image.Point{x, y}.In(p.Rect)) {
		return YUV{}
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	return YUV{Y: s[0], U: s[1], V: s[2], A: s[3]}
}

func (p *YUVAImage) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*4
}

func (p *YUVAImage) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := YUVModel.Convert(c).(YUV)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = c1.Y
	s[1] = c1.U
	s[2] = c1.V
	s[3] = c1.A
}

func (p *YUVAImage) SetYUVA(x, y int, c YUV) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	s := p.Pix[i : i+4 : i+4] // Small cap improves performance, see https://golang.org/issue/27857
	s[0] = c.Y
	s[1] = c.U
	s[2] = c.V
	s[3] = c.A
}

func (p *YUVAImage) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to
	// be inside either r1 or r2 if the intersection is empty. Without
	// explicitly checking for this, the Pix[i:] expression below can
	// panic.
	if r.Empty() {
		return &YUVAImage{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &YUVAImage{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

func NewYUVA(r image.Rectangle) *YUVAImage {
	w, h := r.Dx(), r.Dy()
	pix := make([]uint8, 4*w*h)
	return &YUVAImage{pix, 4 * w, r}
}
