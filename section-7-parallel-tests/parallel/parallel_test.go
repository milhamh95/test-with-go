package parallel

import (
	"fmt"
	"testing"
	"time"
)

// func TestSomething(t *testing.T) {
// 	t.Parallel()
// }

// func TestA(t *testing.T) {
// 	t.Parallel()
// }

func TestB(t *testing.T) {
	fmt.Println("setup")
	defer fmt.Println("deferred teardown")

	// make sure test run first
	t.Run("group", func(t *testing.T) {
		t.Run("sub 1", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("sub1 done")
		})
		t.Run("sub 2", func(t *testing.T) {
			t.Parallel()
			time.Sleep(time.Second)
			fmt.Println("sub2 done")
		})
	})

	fmt.Println("teardown")
}
