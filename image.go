package datagen

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/draw"
	"image/png"
)

// ImageData image base64 data
// len(n) == 0  default width/heght = imageDefaultSize
// len(n) == 1 widht == height == n[0]
// len(n) >= 2 width = n[0] height = n[1]
func ImageData(n ...int) string {
	w, h := getImageSize(n...)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(
		img,
		img.Rect,
		image.NewUniform(randColor()),
		image.Point{},
		draw.Over,
	)
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		return ""
	}
	return "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
}

// ImageURL url image
// len(n) == 0  default width/heght = imageDefaultSize
// len(n) == 1 widht == height == n[0]
// len(n) >= 2 width = n[0] height = n[1]
func ImageURL(n ...int) string {
	w, h := getImageSize(n...)
	c := randColor()
	const tempURL = "https://dummyimage.com/%dx%d/%02x%02x%02x"
	return fmt.Sprintf(
		tempURL,
		w, h, c.R, c.G, c.B,
	)
}

const imageDefaultSize = 128
const imageMaxSize = 2048

func getImageSize(n ...int) (int, int) {
	if len(n) == 0 {
		return imageDefaultSize, imageDefaultSize
	}
	w, h := n[0], n[len(n)-1]
	if w <= 0 {
		w = imageDefaultSize
	}
	if w > imageMaxSize {
		w = imageMaxSize
	}
	if h <= 0 {
		h = imageDefaultSize
	}
	if h > imageMaxSize {
		h = imageMaxSize
	}
	return w, h
}
