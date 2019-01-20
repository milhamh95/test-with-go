package parallel

import (
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	t.Parallel()
}

func TestA(t *testing.T) {
	t.Parallel()
}

func TestB(t *testing.T) {
	t.Parallel()
	t.Run("sub 1", func(t *testing.T) {
		t.Parallel()
		time.Sleep(time.Second)
	})
	t.Run("sub 2", func(t *testing.T) {
		t.Parallel()
		time.Sleep(time.Second * 2)
	})
}
