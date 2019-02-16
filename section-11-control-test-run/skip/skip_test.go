package skip

import "testing"

// var shouldBeSkipped = false

func TestThing(t *testing.T) {
	// if shouldBeSkipped {
	// 	t.Skip()
	// }
	if testing.Short() {
		t.Skip()
	}
	t.Log("this test ran!")
}
