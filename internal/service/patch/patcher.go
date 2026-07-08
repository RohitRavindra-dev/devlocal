package patch

import (
	"fmt"
	"strings"

	"github.com/RohitRavindra-dev/devlocal/internal/filesystem"
	"github.com/RohitRavindra-dev/devlocal/internal/git"
)

func listFilesToPatch(path string, isDir bool) ([]string, error) {

	filesToPatch := []string{}

	if isDir {
		filesWithChanges, err := git.ListFilesWithChangesInDir(path)
		if err != nil {
			return nil, err
		}
		filesToPatch = append(filesToPatch, filesWithChanges...)
	} else {
		hasChanges, err := git.FileHasChanges(path)
		if err != nil {
			return nil, err
		}
		if hasChanges {
			filesToPatch = append(filesToPatch, path)
		}
	}

	return filesToPatch, nil

}

func Run(path string) error {

	if path == "" {
		return fmt.Errorf("[Error] path is empty!")
	}

	isDir, err := filesystem.PathExists(path)
	if err != nil {
		return err
	}

	filesToPatch, err := listFilesToPatch(path, isDir)
	if err != nil {
		return err
	}
	if len(filesToPatch) == 0 {
		return fmt.Errorf("[Error] no files with changes at indicated path: %s", path)
	}
	fmt.Printf("[Started] applying patch to files: \n\t\t- %s\n", strings.Join(filesToPatch, "\n\t\t- "))
	fmt.Println("[Completed] recording new patch into devlocal")
	return nil
}
