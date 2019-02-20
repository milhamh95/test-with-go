package main

import (
	"image"
	"image/draw"
	"image/png"
	"os"
	twgdraw "test-with-go/section-13-external-internal-testing/draw"
)

func main() {
	const w, h = 1000, 1000
	var im draw.Image

	im = image.NewRGBA(image.Rectangle{
		Max: image.Point{
			X: w,
			Y: h,
		},
	})

	im = twgdraw.FibGradient(im)
	f, err := os.Create("image.png")
	if err != nil {
		panic(err)
	}

	err = png.Encode(f, im)
	if err != nil {
		panic(err)
	}
}
