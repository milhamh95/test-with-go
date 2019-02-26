package sub

import "os/exec"

func GitStatus() (string, error) {
	cmd := exec.Command("git", "status")
	status, err := cmd.Output()
	return string(status), err
}
