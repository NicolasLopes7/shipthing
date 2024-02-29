package github

import (
	"os/exec"
)

func Clone(url string, output string) (string, error) {
	out, err := exec.Command("git", "clone", url, output).Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}
