package timing

import (
	"errors"
	"time"
)

var (
	timeoutMultiplier time.Duration = 1
)

var (
	ErrTimeout = errors.New("timing: operation timed out")
	timeAfter  = time.After
)

func DoThing() error {
	done := make(chan bool)
	go func() {
		// do some work
		// in the real code, this would be an interaction with an external API or database
		time.Sleep(500 * time.Millisecond)
		done <- true
	}()

	select {
	case <-done:
		return nil
	case <-timeAfter(100 * time.Millisecond * timeoutMultiplier):
		return ErrTimeout
	}
}
