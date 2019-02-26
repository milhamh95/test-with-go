package golden

import (
	"image/color"
	"image/draw"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func FibGradient(im draw.Image) draw.Image {
	min, max := im.Bounds().Min, im.Bounds().Max
	for x := min.X; x < max.X; x++ {
		for y := min.Y; y < max.Y; y++ {
			im.Set(x, y, color.RGBA{
				R: uint8(fib(x) % 256),
				G: uint8(fib(y) % 256),
				B: uint8(fib(x+y) % 256),
				A: uint8((x + y) % 256),
			})
		}
	}
	return im
}
