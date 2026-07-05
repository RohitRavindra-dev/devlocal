package apply

import (
	"fmt"
	"strings"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/git"
)

func applyPatches(patchFiles []string) error {

	if err := filesystem.ValidatePatchesSetup(); err != nil {
		return err
	}

	fmt.Println("[Running] patch for files: ", strings.Join(patchFiles, ", "))

	return nil
}

func applyOverlook(overlookFiles []string) error {
	fmt.Println("[Running] git skip worktree for files: ", strings.Join(overlookFiles, ", "))
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

	// apply patches
	if patchingErr := applyPatches(config.Patches); patchingErr != nil {
		return patchingErr
	}

	// overlook files
	if overlookErr := applyOverlook(config.Overlook); overlookErr != nil {
		return overlookErr
	}

	return nil
}
