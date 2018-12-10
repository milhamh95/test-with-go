package signal

import "testing"

func TestHandler(t *testing.T) {
	t.Fatal("this is error")
	t.Logf("this is a log with : %v", "ok")

}
