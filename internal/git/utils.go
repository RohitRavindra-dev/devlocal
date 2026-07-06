package git

import (
	"fmt"
	"os/exec"
)

func executeGitCommand(args []string) error {
	cmd := exec.Command("git", args...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("Failed to run git command: %s", out)
	}

	return nil
}

func isFileTracked(file string) bool {
	trackCheckArgs := []string{
		"ls-files",
		"--error-unmatch",
		file,
	}

	err := executeGitCommand(trackCheckArgs)

	return err == nil
}
