package draw_test

import (
	"fmt"
	"image/draw"
	td "test-with-go/section-13-external-internal-testing/draw"
	"testing"
)

func TestFibGradient(t *testing.T) {
	var im draw.Image
	td.FibGradient(im)
}

func TestFibFunc(t *testing.T) {
	fmt.Println(td.A)
	got := td.Fib(2)
	if got != 1 {
		t.Errorf("Fib(2) = %d, want 1", got)
	}
}

// func TestInfo(t *testing.T) {
// 	d := td.Dog{}
// 	td.Info(d)
// }
