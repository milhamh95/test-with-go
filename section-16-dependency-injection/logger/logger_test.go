package logger_test

import (
	"fmt"
	"log"
	"strings"
	logger "test-with-go/section-16-dependency-injection/logger"
	"testing"
)

type fakeLogger struct {
	sb strings.Builder
}

func (fl *fakeLogger) Println(v ...interface{}) {
	fmt.Fprintln(&fl.sb, v...)
}

func (fl *fakeLogger) Printf(format string, v ...interface{}) {
	fmt.Fprintf(&fl.sb, format, v...)
}

func TestDemoV3(t *testing.T) {
	var fl fakeLogger
	logger.DemoV3(fl.Println)

	want := "error in doTheThing():"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV4(t *testing.T) {
	var fl fakeLogger
	logger.DemoV4(&fl)

	want := "error in doTheThing():"
	got := fl.sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}

func TestDemoV5(t *testing.T) {
	var sb strings.Builder
	testLogger := log.New(&sb, "", 0)
	thing := logger.Thing{
		Logger: testLogger,
	}
	thing.DemoV5()
	want := "error in doTheThing():"
	got := sb.String()
	if !strings.Contains(got, want) {
		t.Errorf("Logs = %q; want substring %q", got, want)
	}
}
