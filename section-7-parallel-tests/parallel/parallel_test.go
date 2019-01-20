package parallel

import (
	"fmt"
	"testing"
)

// func TestSomething(t *testing.T) {
// 	t.Parallel()
// }

// func TestA(t *testing.T) {
// 	t.Parallel()
// }

// func TestB(t *testing.T) {
// 	fmt.Println("setup")
// 	defer fmt.Println("deferred teardown")

// 	// make sure test run first
// 	t.Run("group", func(t *testing.T) {
// 		t.Run("sub 1", func(t *testing.T) {
// 			t.Parallel()
// 			time.Sleep(time.Second)
// 			fmt.Println("sub1 done")
// 		})
// 		t.Run("sub 2", func(t *testing.T) {
// 			t.Parallel()
// 			time.Sleep(time.Second)
// 			fmt.Println("sub2 done")
// 		})
// 	})

// 	fmt.Println("teardown")
// }

type TestCase struct {
	arg  int
	want int
}

func TestGotcha(t *testing.T) {
	testCases := []TestCase{
		{2, 5},
		{3, 9},
		{4, 16},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with arg = %d, want = %d", tc.arg, tc.want)
			if tc.arg*tc.arg != tc.want {
				t.Errorf("%d^2 != %d", tc.arg, tc.want)
			}
		})
	}
}
