package git

import (
	"fmt"
	"os/exec"
)

func executeGitCommand(args []string) error {
	cmd := exec.Command("git", args...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("[Error] Failed to run Skip Worktree: %s", out)
	}

	return nil
}

func SkipWorkTree(files []string) error {

	filesToSkip := sanitizeToTrackedFiles(files)
	if len(filesToSkip) == 0 {
		return fmt.Errorf("[Error] No files listed in config are being tracked by git")
	}
	skipWorktreeArgs := append(
		[]string{
			"update-index",
			"--skip-worktree",
		},
		filesToSkip...,
	)
	return executeGitCommand(skipWorktreeArgs)
}

func NoSkipWorkTree(files []string) error {

	filesToRevert := sanitizeToTrackedFiles(files)
	if len(filesToRevert) == 0 {
		return fmt.Errorf("[Error] No files listed in config are being tracked by git")
	}
	skipWorktreeArgs := append(
		[]string{
			"update-index",
			"--no-skip-worktree",
		},
		filesToRevert...,
	)
	return executeGitCommand(skipWorktreeArgs)
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

func sanitizeToTrackedFiles(files []string) []string {

	filesTracked := []string{}

	for _, file := range files {
		if isFileTracked(file) {
			filesTracked = append(filesTracked, file)
		}
	}

	return filesTracked

}

func IsSkipWorktree(file string) (bool, error) {
	out, err := exec.Command(
		"git",
		"ls-files",
		"-v",
		file,
	).CombinedOutput()

	if err != nil {
		return false, err
	}

	return len(out) > 0 && out[0] == 'S', nil
}
