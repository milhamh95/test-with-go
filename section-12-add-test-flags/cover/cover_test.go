package cover

import "testing"

// func TestTriangle(t *testing.T) {
// 	tests := []struct {
// 		base, height, want float64
// 	}{
// 		{2, 5, 5},
// 		{2, 2, 2},
// 		{11, 4, 22},
// 	}

// 	for _, tc := range tests {
// 		got := Triangle(tc.base, tc.height)
// 		if got != tc.want {
// 			t.Errorf("Triangle(%f, %f); want %f", tc.base, tc.height, tc.want)
// 		}
// 	}

// }

// Comment this code if Triangle function is commented
func TestSquare(t *testing.T) {
	for i := 0.0; i < 100.0; i++ {
		want := Triangle(i, i) * 2.0
		got := Square(i)
		if got != want {
			t.Errorf("Square(%f) = %f; want %f", i, got, want)
		}
	}
}

// func TestSquare(t *testing.T) {
// 	for i := 0.0; i < 100.0; i++ {
// 		want := i * i
// 		got := Square(i)
// 		if got != want {
// 			t.Errorf("Square(%f) = %f; want %f", i, got, want)
// 		}
// 	}
// }

func TestDoStuff(t *testing.T) {
	got := DoStuff(1, 1)
	want := 2
	if got != want {
		t.Errorf("DoStuff(5,10) = %d; want = %d", got, want)
	}
}
