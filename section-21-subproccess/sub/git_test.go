package sub_test

import (
	"os/exec"
	"strings"
	"test-with-go/section-21-subproccess/sub"
	"testing"
)

var (
	testHashGit bool
)

func init() {
	testHashGit = false
	if _, err := exec.LookPath("git"); err == nil {
		testHashGit = true
	}

}

func TestGitStatus(t *testing.T) {

	if !testHashGit {
		t.Skip("git not found")
	}

	want := "On branch master"
	got, err := sub.GitStatus()
	if err != nil {
		t.Fatalf("GitStatus() err = %s; want nil", err)
	}
	if !strings.HasPrefix(got, want) {
		t.Errorf("GitStatus() = %s; want prefix %s", got, want)
	}
}
