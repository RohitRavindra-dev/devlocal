package apply

import (
	"fmt"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/git"
)

func applyOverlook(overlookFiles []string) error {
	fmt.Println("[Running] git skip worktree for files: ", overlookFiles)
	if len(overlookFiles) == 0 {
		fmt.Println("[Warn] No files found in overlook section of devlocal config, skipping")
		return nil
	}

	if err := git.SkipWorkTree(overlookFiles); err != nil {
		return err
	}

	fmt.Println("[Completed] git skip worktree")

	return nil

}

func Run() error {
	config, err := filesystem.LoadDevlocalConfig()

	if err != nil {
		return err
	}

	// overlook files
	if overlookErr := applyOverlook(config.Overlook); overlookErr != nil {
		return overlookErr
	}

	// TODO: apply patches

	return nil
}
