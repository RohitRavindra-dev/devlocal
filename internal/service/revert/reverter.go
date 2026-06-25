package revert

import (
	"fmt"
	"strings"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/git"
)

func revertOverlook(overlookedFiles []string) error {
	fmt.Println("[Running] revert git skip worktree for files: ", strings.Join(overlookedFiles, ", "))

	if len(overlookedFiles) == 0 {
		fmt.Println("[Warn] No files found in overlook section of devlocal config, skipping")
		return nil
	}

	if err := git.NoSkipWorkTree(overlookedFiles); err != nil {
		return err
	}

	fmt.Println("[Completed] revert git skip worktree")

	return nil

}

func Run() error {

	config, err := filesystem.LoadDevlocalConfig()

	if err != nil {
		return err
	}

	// revert overlooked files
	if overlookRevertErr := revertOverlook(config.Overlook); overlookRevertErr != nil {
		return overlookRevertErr
	}

	//TODO: revert patches

	return nil
}
