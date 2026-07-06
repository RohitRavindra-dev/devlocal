package revert

import (
	"fmt"
	"strings"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/git"
)

func revertPatches(patchedfiles []string) error {
	if err := filesystem.ValidatePatchesSetup(); err != nil {
		return err
	}

	fmt.Println("[Running] revert patches for files: ", strings.Join(patchedfiles, ", "))

	if len(patchedfiles) == 0 {
		fmt.Println("[Warn] No patch files found in patches section of devlocal config, skipping")
		return nil
	}

	if err := git.RevertPatches(patchedfiles); err != nil {
		return err
	}

	fmt.Println("[Completed] reverting patches")
	return nil
}

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

	//revert patches applied
	if patchesRevertErr := revertPatches(config.Patches); patchesRevertErr != nil {
		return patchesRevertErr
	}

	// revert overlooked files
	if overlookRevertErr := revertOverlook(config.Overlook); overlookRevertErr != nil {
		return overlookRevertErr
	}

	fmt.Println("[Completed] reverting devlocal changes")
	return nil
}
