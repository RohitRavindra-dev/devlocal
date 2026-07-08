package git

import (
	"fmt"
	"os/exec"
	"strings"
)

func executeGitCommand(args []string) error {
	cmd := exec.Command("git", args...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("Failed to run git command: %s", out)
	}

	return nil
}

// TODO make this the only git command util
func executeGitCommandWithOutput(args []string) (string, error) {
	cmd := exec.Command("git", args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return string(out), fmt.Errorf("failed to run git %v: %w\n%s", args, err, out)
	}

	return string(out), nil
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

func FileHasChanges(filePath string) (bool, error) {
	out, err := executeGitCommandWithOutput([]string{
		"diff",
		"--name-only",
		"--",
		filePath,
	})
	if err != nil {
		return false, err
	}

	return strings.TrimSpace(out) != "", nil
}

func ListFilesWithChangesInDir(dirpath string) ([]string, error) {
	filesWithChanges, err := executeGitCommandWithOutput([]string{
		"diff",
		"--name-only",
		"--",
		dirpath,
	})

	if err != nil {
		return nil, err
	}

	filesWithChanges = strings.TrimSpace(filesWithChanges)

	if filesWithChanges == "" {
		return []string{}, nil
	}

	return strings.Split(filesWithChanges, "\n"), nil
}
