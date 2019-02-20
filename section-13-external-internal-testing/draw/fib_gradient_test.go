package draw_test

import (
	"image/draw"
	td "test-with-go/section-13-external-internal-testing/draw"
	"testing"
)

func TestFibGradient(t *testing.T) {
	var im draw.Image
	td.FibGradient(im)
}
