package random

import (
	"math/rand"
	"testing"
	"time"
)

func TestPick(t *testing.T) {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	args := make([]int, 25)

	for i := 0; i < 25; i++ {
		args[i] = r.Int()
	}
	got := Pick(args)
	// if got != args[0] {
	// 	t.Errorf("Pick(%v) = %d; want %d", args, got, args[0])
	// }
	for _, v := range args {
		if got == v {
			return
		}
	}

	t.Errorf("Pick(seed=%v) = %d; not in slice", seed, got)

}
